package main

import (
	"fmt"
	"time"

	"github.com/zeromq/goczmq"
)

func main() {
	fmt.Println("Publishing updates at weather server...")

	publisher, _ := goczmq.NewPub("tcp://*:5557")
	collector, _ := goczmq.NewPull("tcp://*:5558")

	defer publisher.Destroy()
	defer collector.Destroy()

	for {
		message, _ := collector.RecvMessage()
		if len(message) > 0 {
			fmt.Println("I: publishing update", string(message[0]))
			publisher.SendMessage(message)
		}
		time.Sleep(time.Second / 2)
	}
}
