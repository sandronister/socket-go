package abstract

import (
	"github.com/sandronister/socket-go/pkg/devices/repositories"
	"github.com/sandronister/socket-go/pkg/devices/service"
	"gitlab.com/gobrax-dev/gobrax-tool/broker/types"
)

type AbstractUseCase struct {
	DeviceRepo   repositories.IDeviceRepository
	Broker       types.IBroker
	BlacklistSVC service.IBlackListService
}
