package uruptela

import (
	"github.com/sandronister/socket-go/pkg/devices"
	customerrors "github.com/sandronister/socket-go/pkg/devices/custom_errors"
)

func (u *UseRuptela) ProcessBehave(device devices.IDevice) (devices.ProtocolBehave, *customerrors.Error) {
	cmd := device.ProcessHeader(u.BlacklistSVC)
	return u.getParameter(device, cmd)
}
