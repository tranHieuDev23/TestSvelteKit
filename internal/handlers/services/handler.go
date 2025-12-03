package services

import (
	"context"
	"fmt"
	"hello-world/internal/handlers/models"
	"hello-world/internal/handlers/services/rpcserver"
)

type Handler struct{}

func NewHandler() rpcserver.ServiceServer {
	return &Handler{}
}

func (h Handler) GetGreeting(ctx context.Context, in *models.GetGreetingRequest) (*models.GetGreetingResponse, error) {
	return &models.GetGreetingResponse{
		Greeting: fmt.Sprintf("Hello, %s!", in.Thing),
	}, nil
}
