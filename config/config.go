package config

import (
	"context"
	"log"

	"cloud.google.com/go/pubsub"
)

func CreatePubSubClient(projectId string) *pubsub.Client {
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectId)
	if err != nil {
		log.Fatalf("Failed to create Pub/Sub client: %v", err)
	}
	return client
}
