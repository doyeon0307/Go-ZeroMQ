package main

import (
	"fmt"
	"time"

	"github.com/zeromq/goczmq"
)

func main() {
	context, _ := goczmq.NewRep("tcp://*:5555")

	for {
		message, _ := context.RecvMessage()
		fmt.Printf("Received request: %s\n", string(message[0]))

		time.Sleep(1 * time.Second)

		context.SendMessage([][]byte{[]byte("World")})
	}
}
