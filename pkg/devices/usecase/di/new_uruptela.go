package di

import (
	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/socket-go/pkg/devices/database/connection/postgre"
	"github.com/sandronister/socket-go/pkg/devices/database/repository"
	"github.com/sandronister/socket-go/pkg/devices/service/blacklist"
	"github.com/sandronister/socket-go/pkg/devices/usecase"
	"github.com/sandronister/socket-go/pkg/devices/usecase/uruptela"
)

func NewRuptelaUsecase() (usecase.IDeviceUseCase, error) {
	conn, err := postgre.GetPostGresConnection()
	if err != nil {
		return nil, err
	}

	broker := factory.GetBroker()
	repo := repository.NewDeviceRepository(conn)
	blacklistService := blacklist.Factory(repo)
	return uruptela.NewRuptelaUsecase(repo, blacklistService, broker), nil
}
