package di

import (
	"github.com/sandronister/socket-go/pkg/devices/database/connection/postgre"
	"github.com/sandronister/socket-go/pkg/devices/database/repository"
	"github.com/sandronister/socket-go/pkg/devices/service/blacklist"
	"github.com/sandronister/socket-go/pkg/devices/usecase"
	"github.com/sandronister/socket-go/pkg/devices/usecase/uruptela"
	"gitlab.com/gobrax-dev/gobrax-tool/broker/factory"
	cloudFactory "gitlab.com/gobrax-dev/gobrax-tool/cloud/factory"
)

func NewRuptelaUsecase() (usecase.IDeviceUseCase, error) {
	conn, err := postgre.GetPostGresConnection()
	if err != nil {
		return nil, err
	}

	cloud, err := cloudFactory.GetCloudInstance()

	if err != nil {
		return nil, err
	}

	broker := factory.GetBroker()
	repo := repository.NewDeviceRepository(conn)
	blacklistService := blacklist.Factory(repo)
	return uruptela.NewRuptelaUsecase(repo, blacklistService, broker, cloud), nil
}
