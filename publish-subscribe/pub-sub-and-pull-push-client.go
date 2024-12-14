package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zeromq/goczmq"
)

func main() {
	subscriber, _ := goczmq.NewSub("tcp://localhost:5557", "")
	defer subscriber.Destroy()

	publisher, _ := goczmq.NewPush("tcp://localhost:5558")
	defer publisher.Destroy()

	rand.Seed(time.Now().UnixNano())
	for {
		msg, err := subscriber.RecvMessageNoWait()
		if err == nil {
			fmt.Printf("I: received message %s\n", string(msg[0]))
		} else {
			randNum := rand.Intn(100) + 1
			if randNum < 10 {
				publisher.SendMessage([][]byte{[]byte(fmt.Sprintf("%d", randNum))})
				fmt.Printf("I: sending message %d\n", randNum)
			}
		}
		time.Sleep(time.Second / 2)
	}
}
