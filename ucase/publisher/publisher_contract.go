package publisher

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/fajarcandraaa/go-message-broker/entity"
)

type Publisher interface {
	Publish(ctx context.Context, message entity.Message) *pubsub.PublishResult
}
