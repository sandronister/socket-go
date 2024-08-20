package di

import (
	"github.com/sandronister/socket-go/internal/infra/handler"
	"github.com/sandronister/socket-go/internal/infra/web"
	"github.com/sandronister/socket-go/pkg/devices/usecase/di"
)

func NewServer() (*web.Server, error) {
	usecase, err := di.NewRuptelaUsecase()
	if err != nil {
		return nil, err
	}
	handler := handler.NewTcpHandler(usecase, 3)
	return web.NewServer(handler), nil
}
