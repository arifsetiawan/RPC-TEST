package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"zombiezen.com/go/capnproto2/rpc"

	gr "./../greeter"
)

const (
	address = ":60051"
)

// greeter is a local implementation of Greeter.
type greeter struct{}

func (g greeter) SayHello(call gr.Greeter_sayHello) error {
	// read incoming parameter
	name, err := call.Params.Name()
	if err != nil {
		return err
	}

	//return Hello + Name
	return call.Results.SetRep("Hello " + name)
}

func handleRequest(c net.Conn) {
	// Create a new locally implemented Greeter.
	// log.Println("Instantiating Greeter object...")
	main := gr.Greeter_ServerToClient(greeter{})

	// Listen for calls, using the HashFactory as the bootstrap interface.
	// log.Println("Opening RPC connection...")
	conn := rpc.NewConn(rpc.StreamTransport(c), rpc.MainInterface(main.Client))
	// Wait for connection to abort.
	// log.Println("Waiting for connection to abort...")
	_ = conn.Wait()
	// if err != nil {
	// 	log.Printf("failed to open RPC connection: %v", err)
	// }
}

func main() {
	// Set up a Listener.
	// log.Printf("Listening to port %s...", address)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("could not create Listener: %v", err)
	}
	defer ln.Close()

	// log.Println("Acceptimg RPC connection...")
	// c, err := ln.Accept()
	// if err != nil {
	// 	fmt.Println("Error accepting: ", err.Error())
	// 	os.Exit(1)
	// }
	// // Handle connections in a new goroutine.
	// handleRequest(c)

	for {
		// Listen for an incoming connection.
		// log.Println("Acceptimg RPC connection...")
		c, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(c)
	}
}
