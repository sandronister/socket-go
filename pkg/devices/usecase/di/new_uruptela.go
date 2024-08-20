package di

import (
	"fmt"

	"github.com/sandronister/go_broker/pkg/broker/factory"
	"github.com/sandronister/socket-go/pkg/devices/usecase"
	"github.com/sandronister/socket-go/pkg/devices/usecase/uruptela"
)

func NewRuptelaUsecase() (usecase.IUseCase, error) {
	//urlParse, err := pg.ParseURL(os.Getenv("DATABASE_URL"))
	//
	fmt.Println("Connected to database")
	//conn := pg.Connect(urlParse)

	broker := factory.GetBroker()
	//repo := repository.NewDeviceRepository(conn)
	return uruptela.NewRuptelaUsecase(nil, broker), nil
}
