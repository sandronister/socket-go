package di

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/socket-go/internal/infra/database/postgre/repositories"
	"github.com/sandronister/socket-go/internal/infra/handler"
	"github.com/sandronister/socket-go/internal/infra/web"
	"github.com/sandronister/socket-go/internal/usecase"
)

func NewServer() (*web.Server, error) {
	conn, err := pgxpool.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	broker := factory.GetBroker()
	repo := repositories.NewDeviceRepository(conn)
	usecase := usecase.NewRuptelaUsecase(repo, broker)
	handler := handler.NewTcpHandler(usecase, 3)
	return web.NewServer(handler), nil
}
