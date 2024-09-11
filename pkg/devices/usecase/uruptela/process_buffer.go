package uruptela

import (
	"github.com/sandronister/socket-go/pkg/devices"
	customerrors "github.com/sandronister/socket-go/pkg/devices/custom_errors"
	"github.com/sandronister/socket-go/pkg/devices/entities/ruptela"
)

func (u *UseRuptela) ProcessBehave(device devices.IDevice) (devices.ProtocolBehave, *customerrors.Error) {
	cmd := device.ProcessHeader(u.BlacklistSVC)

	if cmd != ruptela.COMMAND_GOBRAX_BLOCKED_IMEI {
		err := u.Dispatch(device)
		if err != nil {
			return devices.BEHAVE_REPLY_AND_CLOSE_CONN, err
		}
	}

	return u.getParameter(device, cmd)
}
