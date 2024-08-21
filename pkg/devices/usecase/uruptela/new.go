package uruptela

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/socket-go/pkg/devices/repositories"
	"github.com/sandronister/socket-go/pkg/devices/service"
	"github.com/sandronister/socket-go/pkg/devices/usecase"
	"github.com/sandronister/socket-go/pkg/devices/usecase/abstract"
)

func NewRuptelaUsecase(deviceRepo repositories.IDeviceRepository, blacklist service.IBlackListService, broker types.IBroker) usecase.IDeviceUseCase {
	r := &UseRuptela{
		AbstractUseCase: abstract.AbstractUseCase{
			DeviceRepo:   deviceRepo,
			Broker:       broker,
			BlacklistSVC: blacklist,
		},
	}

	return r
}
