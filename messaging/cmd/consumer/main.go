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
		println("Cannot connect")
		fmt.Printf("error: %v+", err)
		os.Exit(1)
	}
	defer nc.Close()
	ch := make(chan *nats.Msg, 64)
	_, err = nc.ChanSubscribe(natsTopic, ch)
	if err != nil {
		println("Cannot connect")
		os.Exit(1)
	}
	println("Waiting for messages:")
	for {
		msg := <-ch
		println(string(msg.Data))
	}

}
