package uruptela

import (
	"github.com/sandronister/socket-go/pkg/devices/repositories"
	"github.com/sandronister/socket-go/pkg/devices/service"
	"github.com/sandronister/socket-go/pkg/devices/usecase"
	"github.com/sandronister/socket-go/pkg/devices/usecase/abstract"
	"gitlab.com/gobrax-dev/gobrax-tool/broker/types"
	tcloud "gitlab.com/gobrax-dev/gobrax-tool/cloud/types"
)

func NewRuptelaUsecase(deviceRepo repositories.IDeviceRepository, blacklist service.IBlackListService, broker types.IBroker, cloud tcloud.ICloudProvider) usecase.IDeviceUseCase {
	r := &UseRuptela{
		AbstractUseCase: abstract.AbstractUseCase{
			DeviceRepo:   deviceRepo,
			Broker:       broker,
			BlacklistSVC: blacklist,
		},
		cloud: cloud,
	}

	return r
}
