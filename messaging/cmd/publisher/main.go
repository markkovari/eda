package main

import (
	"fmt"
	"os"

	"github.com/nats-io/nats.go"
)

const topic string = "the_message_topic"

func main() {
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		println("Cannot connect")
		os.Exit(1)
	}
	defer nc.Close()
	currentMessage := 0

	for i := 0; i < 100; i++ {
		nextMessage := "Sending message " + fmt.Sprint(currentMessage)
		nc.Publish(topic, []byte(nextMessage))
		currentMessage += 1
		println("Successfully sent")
	}
}
