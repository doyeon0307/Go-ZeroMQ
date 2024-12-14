package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/zeromq/goczmq"
)

func main() {
	fmt.Println("Publishing updates at weather server...")

	publisher, _ := goczmq.NewPub("tcp://*:5556")
	defer publisher.Destroy()

	rand.Seed(time.Now().UnixNano())

	for {
		zipcode := rand.Intn(99999) + 1
		temperature := rand.Intn(215) - 80
		relhumidity := rand.Intn(50) + 10

		msg := fmt.Sprintf("%d %d %d", zipcode, temperature, relhumidity)
		publisher.SendMessage([][]byte{[]byte(msg)})

		fmt.Println("Sent:", msg)
		time.Sleep(time.Second / 4)
	}
}
