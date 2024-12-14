package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/zeromq/goczmq"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run client.go client_id")
		return
	}
	clientID := os.Args[1]

	subscriber, _ := goczmq.NewSub("tcp://localhost:5557", "")
	publisher, _ := goczmq.NewPush("tcp://localhost:5558")

	defer subscriber.Destroy()
	defer publisher.Destroy()

	rand.Seed(time.Now().UnixNano())

	for {
		message, err := subscriber.RecvMessageNoWait()
		if err == nil && len(message) > 0 {
			fmt.Printf("%s: receive status => %s\n", clientID, string(message[0]))
		} else {
			randNum := rand.Intn(100) + 1
			if randNum < 10 {
				time.Sleep(time.Second)
				msg := fmt.Sprintf("(%s:ON)", clientID)
				publisher.SendMessage([][]byte{[]byte(msg)})
				fmt.Printf("%s: send status - activated\n", clientID)
			} else if randNum > 90 {
				time.Sleep(time.Second)
				msg := fmt.Sprintf("(%s:OFF)", clientID)
				publisher.SendMessage([][]byte{[]byte(msg)})
				fmt.Printf("%s: send status - deactivated\n", clientID)
			}
		}
		time.Sleep(time.Second / 4)
	}
}
