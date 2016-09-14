package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/context"
	"zombiezen.com/go/capnproto2/rpc"

	gr "./../greeter"
)

const (
	address     = "localhost:60051"
	defaultName = "world"
)

func invoke(g gr.Greeter, ctx context.Context, name string) {

	// Call methods on `g`, and they will be sent over `c`.
	// log.Println("Attempting to call a remote procedure...")
	result, err := g.SayHello(ctx, func(p gr.Greeter_sayHello_Params) error {
		err := p.SetName(name)
		if err != nil {
			log.Fatalf("could not set parameter 'name': %v", err)
		}
		return err
	}).Struct()

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// log.Println("Attempting to get the result from server...")
	_, err = result.Rep()
	if err != nil {
		log.Fatalf("could not get RPC result from server : %v", err)
	}
	// log.Printf("Got the following : %v", r)
}

func syncTest(c gr.Greeter, ctx context.Context, name string) {
	i := 10000
	t := time.Now().UnixNano()
	for ; i > 0; i-- {
		invoke(c, ctx, name)
	}
	fmt.Println((time.Now().UnixNano() - t) / 1000000)
}

func asyncTest(c [20]gr.Greeter, ctx context.Context, name string) {
	var wg sync.WaitGroup
	wg.Add(10000)

	i := 10000
	t := time.Now().UnixNano()
	for ; i > 0; i-- {
		go func(counter int) {
			invoke(c[counter%20], ctx, name)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println((time.Now().UnixNano() - t) / 1000000)
}

func main() {
	// Set up a connection to the server.
	// log.Println("Opening TCP connection...")
	c, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// Create a connection that we can use to get the Greeter.
	// log.Println("Opening RPC connection...")
	conn := rpc.NewConn(rpc.StreamTransport(c))
	defer conn.Close()
	var g [20]gr.Greeter

	// Contact the server and print out its response.
	name := defaultName
	sync := true
	if len(os.Args) > 1 {
		sync, err = strconv.ParseBool(os.Args[1])
	}
	if err != nil {
		log.Println("Error parsing argument!")
	}

	ctx := context.Background()

	//warm up
	i := 0
	// log.Println("Setting up 20 Greeter clients...")
	for ; i < 20; i++ {
		// Get the "bootstrap" interface.  This is the capability set with
		// rpc.MainInterface on the remote side.
		g[i] = gr.Greeter{Client: conn.Bootstrap(ctx)}
		invoke(g[i], ctx, name)
	}

	// log.Println("Begin test...")
	if sync {
		syncTest(g[0], ctx, name)
	} else {
		asyncTest(g, ctx, name)
	}
}
