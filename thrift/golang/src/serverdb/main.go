package main

import (
	"fmt"
	"log"
	"os"

	"github.com/couchbase/gocb"

	"git.apache.org/thrift.git/lib/go/thrift"

	"./../greeter"
)

const (
	NetworkAddr = "localhost:9090"
)

var bucket *gocb.Bucket

type GreeterHandler struct {
}

// User is a single user
type User struct {
	Entity     string   `json:"_entity,omitempty"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	PositionID string   `json:"position_id,omitempty"`
	Position   string   `json:"position,omitempty"`
	Group      []string `json:"groups,omitempty"`
}

func NewGreeterHandler() *GreeterHandler {
	return &GreeterHandler{}
}

func (p *GreeterHandler) SayHello(request *greeter.HelloRequest) (r *greeter.HelloReply, err error) {

	listUserQuery := gocb.NewN1qlQuery("SELECT u.email, u.name, u.position_id, u.position FROM gobase u WHERE _entity='user'")
	rows, err := bucket.ExecuteN1qlQuery(listUserQuery, nil)
	if err != nil {
		fmt.Fatal("Failed to connect to database!")
	}
	var user User
	var users []User
	for i := 0; rows.Next(&user); i++ {
		users = append(users, user)
	}
	_ = rows.Close()

	reply := greeter.NewHelloReply()
	reply.Message = "Hello " + request.Message
	return reply, nil
}

func main() {
	cluster, err := gocb.Connect("couchbase://192.168.99.100")
	if err != nil {
		log.Fatal("Failed to connect to database!")
		return
	}
	bucket, err = cluster.OpenBucket(
		"gobase", "Test1234")
	if err != nil {
		log.Fatal("Failed to open bucket!")
		return
	}

	var protocolFactory thrift.TProtocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	var transportFactory thrift.TTransportFactory = thrift.NewTBufferedTransportFactory(8192)
	transport, err := thrift.NewTServerSocket(NetworkAddr)
	if err != nil {
		fmt.Println("Error!", err)
		os.Exit(1)
	}

	handler := NewGreeterHandler()
	processor := greeter.NewGreeterProcessor(handler)
	server := thrift.NewTSimpleServer4(processor, transport, transportFactory, protocolFactory)
	fmt.Println("Starting the simple server... on ", NetworkAddr)
	server.Serve()
}
