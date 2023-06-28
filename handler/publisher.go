package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fajarcandraaa/go-message-broker/entity"
	"github.com/fajarcandraaa/go-message-broker/ucase/publisher"
)

type PublisherHandler struct {
	service publisher.Publisher
}

func NewPublisherHandler(service publisher.Publisher) *PublisherHandler {
	return &PublisherHandler{
		service: service,
	}
}

func (h *PublisherHandler) PublishHandler(w http.ResponseWriter, r *http.Request) {
	var requestBody entity.Message
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	result := h.service.Publish(context.Background(), requestBody)
	_, err = result.Get(context.Background())
	if err != nil {
		http.Error(w, "Failed to publish message", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, "Message published successfully")
}
