package entity

import (
	"fmt"

	"cloud.google.com/go/pubsub"
)

type Message struct {
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

func MessageHandler(msg *pubsub.Message) {
	fmt.Printf("Received message: %s\n", string(msg.Data))
	msg.Ack()
}
