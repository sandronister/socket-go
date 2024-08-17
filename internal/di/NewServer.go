package di

import (
	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/socket-go/config"
	"github.com/sandronister/socket-go/internal/infra/handler"
	"github.com/sandronister/socket-go/internal/infra/web"
)

func NewServer(conf *config.Conf) *web.Server {
	broker := factory.GetBroker()
	handler := handler.NewTcpHandler(broker, 3)
	return web.NewServer(conf, handler)
}
