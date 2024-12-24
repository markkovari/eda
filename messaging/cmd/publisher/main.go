package main

import (
	"fmt"
	"os"

	"github.com/nats-io/nats.go"
)

var natsUrl = os.Getenv("NATS_URL")
var natsUser = os.Getenv("NATS_USER")
var natsPassword = os.Getenv("NATS_PASSWORD")
var natsTopic = os.Getenv("NATS_TOPIC")

func main() {

	nc, err := nats.Connect(natsUrl, nats.UserInfo(natsUser, natsPassword))

	if err != nil {
		fmt.Print("Cannot connect")
		os.Exit(1)
	}
	defer nc.Close()
	currentMessage := 0

	for i := 0; i < 100; i++ {
		nextMessage := "Sending message " + fmt.Sprint(currentMessage)
		nc.Publish(natsTopic, []byte(nextMessage))
		currentMessage += 1
		fmt.Printf("Successfully sent: %d\n", currentMessage)
	}
}
