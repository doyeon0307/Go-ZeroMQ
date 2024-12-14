package main

import (
	"fmt"

	"github.com/zeromq/goczmq"
)

func main() {
	publisher, _ := goczmq.NewPub("tcp://*:5557")
	collector, _ := goczmq.NewPull("tcp://*:5558")

	defer publisher.Destroy()
	defer collector.Destroy()

	for {
		message, _ := collector.RecvMessage()
		if len(message) > 0 {
			fmt.Println("server: publishing update => ", string(message[0]))
			publisher.SendMessage(message)
		}
	}
}
