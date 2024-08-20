package uruptela

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/socket-go/pkg/devices/repositories"
	"github.com/sandronister/socket-go/pkg/devices/usecase"
	"github.com/sandronister/socket-go/pkg/devices/usecase/abstract"
)

func NewRuptelaUsecase(deviceRepo repositories.IDeviceRepository, broker types.IBroker) usecase.IUseCase {
	r := &usecase_t{
		AbstractUseCase: abstract.AbstractUseCase{
			DeviceRepo: deviceRepo,
			Broker:     broker,
		},
	}

	//r.InitBlackListDevice()

	return r
}
