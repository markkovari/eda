package main

import (
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
	ch := make(chan *nats.Msg, 64)
	_, err = nc.ChanSubscribe(topic, ch)
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
