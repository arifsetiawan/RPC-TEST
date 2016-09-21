package main

import (
	"log"
	"net"

	"github.com/couchbase/gocb"

	pb "./../greeter"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

var bucket *gocb.Bucket

type server struct{}

// User is a single user
type User struct {
	Entity     string   `json:"_entity,omitempty"`
	Name       string   `json:"name"`
	Email      string   `json:"email"`
	PositionID string   `json:"position_id,omitempty"`
	Position   string   `json:"position,omitempty"`
	Group      []string `json:"groups,omitempty"`
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	listUserQuery := gocb.NewN1qlQuery("SELECT u.email, u.name, u.position_id, u.position FROM gobase u WHERE _entity='user'")
	rows, err := bucket.ExecuteN1qlQuery(listUserQuery, nil)
	if err != nil {
		log.Fatal("Failed to connect to database!")
	}
	var user User
	var users []User
	for i := 0; rows.Next(&user); i++ {
		users = append(users, user)
	}
	_ = rows.Close()

	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
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

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	s.Serve(lis)
}
