package di

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/socket-go/config"
	"github.com/sandronister/socket-go/internal/infra/database/postgree/repositories"
	"github.com/sandronister/socket-go/internal/infra/handler"
	"github.com/sandronister/socket-go/internal/infra/web"
	"github.com/sandronister/socket-go/internal/usecase"
)

func NewServer(conf *config.Conf, conn *pgxpool.Pool) *web.Server {
	broker := factory.GetBroker()
	repo := repositories.NewDeviceRepository(conn)
	usecase := usecase.NewRuptelaUsecase(repo, broker)
	handler := handler.NewTcpHandler(usecase, 3)
	return web.NewServer(conf, handler)
}
