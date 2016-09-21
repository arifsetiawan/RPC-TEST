package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"./../greeter"

	"git.apache.org/thrift.git/lib/go/thrift"
)

const (
	address     = "localhost:9090"
	defaultName = "world"
)

func syncTest(client *greeter.GreeterClient, name string) {

	request := greeter.NewHelloRequest()
	request.Message = name

	i := 100
	t := time.Now().UnixNano()
	for ; i > 0; i-- {
		client.SayHello(request)
	}
	//fmt.Println("took", (time.Now().UnixNano()-t)/1000000, "ms")
	fmt.Println((time.Now().UnixNano() - t) / 1000000)
}

func asyncTest(client [20]*greeter.GreeterClient, name string) {

	request := greeter.NewHelloRequest()
	request.Message = name

	var locks [20]sync.Mutex
	var wg sync.WaitGroup
	wg.Add(100)

	i := 100
	t := time.Now().UnixNano()
	for ; i > 0; i-- {
		go func(index int) {
			locks[index%20].Lock()
			client[index%20].SayHello(request)
			wg.Done()
			locks[index%20].Unlock()
		}(i)
	}
	wg.Wait()
	//fmt.Println("took", (time.Now().UnixNano()-t)/1000000, "ms")
	fmt.Println((time.Now().UnixNano() - t) / 1000000)
}

func main() {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	var client [20]*greeter.GreeterClient

	request := greeter.NewHelloRequest()
	request.Message = defaultName

	//warm up
	for i := 0; i < 20; i++ {
		transport, err := thrift.NewTSocket(address)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error resolving address:", err)
			os.Exit(1)
		}
		useTransport := transportFactory.GetTransport(transport)
		defer transport.Close()

		if err := transport.Open(); err != nil {
			fmt.Fprintln(os.Stderr, "Error opening socket to localhost:9090", " ", err)
			os.Exit(1)
		}

		client[i] = greeter.NewGreeterClientFactory(useTransport, protocolFactory)
		client[i].SayHello(request)
	}

	sync := true
	if len(os.Args) > 1 {
		sync, _ = strconv.ParseBool(os.Args[1])
	}

	if sync {
		syncTest(client[0], defaultName)
	} else {
		asyncTest(client, defaultName)
	}
}
