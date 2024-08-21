package di

import (
	"os"

	"github.com/go-pg/pg"
	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/socket-go/pkg/devices/database/repository"
	"github.com/sandronister/socket-go/pkg/devices/service/blacklist"
	"github.com/sandronister/socket-go/pkg/devices/usecase"
	"github.com/sandronister/socket-go/pkg/devices/usecase/uruptela"
)

func NewRuptelaUsecase() (usecase.IDeviceUseCase, error) {
	urlParse, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	conn := pg.Connect(urlParse)

	broker := factory.GetBroker()
	repo := repository.NewDeviceRepository(conn)
	blacklistService := blacklist.Factory(repo)
	return uruptela.NewRuptelaUsecase(repo, blacklistService, broker), nil
}
