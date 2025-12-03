package services

import "hello-world/internal/handlers/models"

// Service description...
//
//go:generate genpjrpc -search.name=Service -print.content.tags_clients=mockgen  -print.content.tags_server=mockgen -print.place.path_swagger_file=../../../api/swagger.json
type Service interface {
	GetGreeting(models.GetGreetingRequest) models.GetGreetingResponse
}
