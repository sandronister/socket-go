package abstract

import (
	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/socket-go/pkg/devices/repositories"
)

type AbstractUseCase struct {
	DeviceRepo      repositories.IDeviceRepository
	Broker          types.IBroker
	BlackListDevice map[string]bool
}
