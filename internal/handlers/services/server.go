package services

import (
	"context"
	"encoding/json"
	"hello-world/internal/handlers/services/rpcserver"
	"log"
	"net/http"

	"gitlab.com/pjrpc/pjrpc/v2"
)

type Server interface {
	Start() error
}

type server struct {
	httpServer *http.Server
}

func NewServer(handler rpcserver.ServiceServer) Server {
	serverHTTP := pjrpc.NewServerHTTP()
	serverHTTP.SetLogger(log.Writer())

	rpcserver.RegisterServiceServer(serverHTTP, handler, pjrpc.Middleware(func(next pjrpc.Handler) pjrpc.Handler {
		return func(ctx context.Context, params json.RawMessage) (any, error) {
			log.Printf("incoming request: %v\n", string(params))
			response, err := next(ctx, params)
			log.Printf("outgoing response: %v\n", response)
			return response, err
		}
	}))

	mux := http.NewServeMux()
	mux.Handle("POST /api", serverHTTP)
	mux.Handle("POST /api/", serverHTTP)

	return &server{
		httpServer: &http.Server{
			Addr:    ":8080",
			Handler: mux,
		},
	}
}

func (s server) Start() error {
	log.Println("starting server")
	return s.httpServer.ListenAndServe()
}
