package main

import (
	"fmt"

	"github.com/zeromq/goczmq"
)

func main() {
	fmt.Println("Connecting to hello world server...")
	context, _ := goczmq.NewReq("tcp://localhost:5555")
	defer context.Destroy()

	for request := 0; request < 10; request++ {
		fmt.Printf("Sending request %d ...\n", request)
		context.SendFrame([]byte("Hello"), 0)

		message, _, _ := context.RecvFrame()
		fmt.Printf("Received reply %d [ %s ]\n", request, string(message))
	}
}
