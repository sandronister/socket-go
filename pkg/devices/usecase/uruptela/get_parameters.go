package uruptela

import (
	"github.com/sandronister/socket-go/pkg/devices"
	customerrors "github.com/sandronister/socket-go/pkg/devices/custom_errors"
	"github.com/sandronister/socket-go/pkg/devices/entities/ruptela"
)

func (u *UseRuptela) getParameter(device devices.IDevice) (devices.ProtocolBehave, *customerrors.Error) {
	cmd := device.ProcessHeader(u.BlacklistSVC)

	switch cmd {
	case ruptela.COMMAND_DYNAMIC_IDENTFIC_PACKET:
		return devices.BEHAVE_REPLY_AND_KEEP_ALIVE, nil
	}
	return devices.BEHAVE_REPLY_AND_CLOSE_CONN, nil
}
