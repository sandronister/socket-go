package web

import (
	"os"

	"github.com/sandronister/socket-go/internal/infra/handler"
)

func NewServer(handler handler.IHandler) *Server {
	return &Server{
		host:    os.Getenv("SOCKET_HOST"),
		port:    os.Getenv("SOCKET_PORT"),
		handler: handler,
	}
}
