package main

import (
	"hello-world/internal/handlers/services"
	"log"
)

func main() {
	handler := services.NewHandler()
	server := services.NewServer(handler)
	if err := server.Start(); err != nil {
		log.Println(err)
	}
}
