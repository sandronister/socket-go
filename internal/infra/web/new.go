package web

import (
	"github.com/sandronister/socket-go/config"
	"github.com/sandronister/socket-go/internal/infra/handler"
)

func NewServer(conf *config.Conf, handler *handler.TcpHandler) *Server {
	return &Server{
		host:      conf.SOCKET_HOST,
		port:      conf.SOCKET_PORT,
		state:     STATE_SET_TIMEOUT,
		nextState: STATE_SET_TIMEOUT,
		handler:   handler,
	}
}
