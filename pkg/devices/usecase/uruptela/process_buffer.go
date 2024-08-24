package uruptela

import (
	"github.com/sandronister/socket-go/pkg/devices"
	customerrors "github.com/sandronister/socket-go/pkg/devices/custom_errors"
	"github.com/sandronister/socket-go/pkg/devices/entities/ruptela"
)

func (u *UseRuptela) ProcessBehave(device devices.IDevice) (devices.ProtocolBehave, *customerrors.Error) {
	cmd := device.ProcessHeader(u.BlacklistSVC)
	if cmd != ruptela.COMMAND_GOBRAX_BLOCKED_IMEI {
		u.SendToBucket(device)
	}
	return u.getParameter(device, cmd)
}
