package handler

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"cloud.google.com/go/pubsub"
	"github.com/fajarcandraaa/go-message-broker/config"
	"github.com/fajarcandraaa/go-message-broker/entity"
)

type SubscriberHandler struct {
	id string
}

func NewSubscriberHandler(id string) *SubscriberHandler {
	return &SubscriberHandler{
		id: id,
	}
}

func (s *SubscriberHandler) Subscriber() {
	subscriber := config.CreatePubSubClient(s.id)
	// subscribtionName := "projects/pubsubx-391205/subscriptions/testpubsubx-sub"
	subscribtionName := "testpubsubx-sub"

	sub := subscriber.Subscription(subscribtionName)
	cctx, cancel := context.WithCancel(context.Background())
	err := sub.Receive(cctx, func(ctx context.Context, msg *pubsub.Message) {
		go entity.MessageHandler(msg)
	})
	if err != nil {
		log.Fatalf("Failed to start Pub/Sub subscriber: %v", err)
	}

	// Wait for termination signal to gracefully shutdown the subscriber.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch

	cancel()
	<-time.After(10 * time.Second)
}
