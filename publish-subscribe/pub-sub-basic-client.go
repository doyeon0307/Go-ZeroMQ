package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/zeromq/goczmq"
)

func main() {
	subscriber, _ := goczmq.NewSub("tcp://localhost:5556", "")
	defer subscriber.Destroy()

	fmt.Println("Collecting updates from weather server...")

	zipFilter := "10001"
	if len(os.Args) > 1 {
		zipFilter = os.Args[1]
	}

	subscriber.SetSubscribe(zipFilter)

	time.Sleep(time.Second)

	fmt.Println("Subscribed to zipcode:", zipFilter)

	totalTemp := 0
	for updateNbr := 0; updateNbr < 20; updateNbr++ {
		msg, _ := subscriber.RecvMessage()
		if len(msg) == 0 {
			continue
		}

		parts := strings.Split(string(msg[0]), " ")
		if len(parts) < 2 {
			continue
		}

		temperature, _ := strconv.Atoi(parts[1])

		totalTemp += temperature
		fmt.Printf("Receive temperature for zipcode '%s' was %d F\n", zipFilter, temperature)
	}

	fmt.Printf("Average temperature for zipcode '%s' was %d F\n",
		zipFilter, totalTemp/20)
}
