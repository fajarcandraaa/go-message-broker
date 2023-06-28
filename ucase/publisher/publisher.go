package publisher

import (
	"context"

	"cloud.google.com/go/pubsub"
	"github.com/fajarcandraaa/go-message-broker/entity"
)

type publihser struct {
	p *pubsub.Client
}

func NewPublisher(p *pubsub.Client) *publihser {
	return &publihser{
		p: p,
	}
}

var _ Publisher = &publihser{}

// Publish implements Publisher.
func (pb *publihser) Publish(ctx context.Context, message entity.Message) *pubsub.PublishResult {
	result := pb.p.Topic(message.Topic).Publish(ctx, &pubsub.Message{
		Data: []byte(message.Message),
	})

	return result
}
