package di

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/socket-go/pkg/devices/database/repository"
	"github.com/sandronister/socket-go/pkg/devices/usecase"
	"github.com/sandronister/socket-go/pkg/devices/usecase/uruptela"
)

func NewRuptelaUsecase() (usecase.IUseCase, error) {
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	broker := factory.GetBroker()
	repo := repository.NewDeviceRepository(conn)
	return uruptela.NewRuptelaUsecase(repo, broker), nil
}
