package uruptela

import (
	"github.com/sandronister/socket-go/pkg/devices"
	customerrors "github.com/sandronister/socket-go/pkg/devices/custom_errors"
	"github.com/sandronister/socket-go/pkg/devices/entities/ruptela"
)

func (u *UseRuptela) getParameter(device devices.IDevice, cmd ruptela.TCommandProtocol) (devices.ProtocolBehave, *customerrors.Error) {

	switch cmd {
	case ruptela.COMMAND_DYNAMIC_IDENTFIC_PACKET:
		return devices.BEHAVE_REPLY_AND_KEEP_ALIVE, device.ProcessDynamic()

	case ruptela.COMMAND_DIAGNOSTIC_TROUBLE_CODES:
		return devices.BEHAVE_REPLY_AND_CLOSE_CONN, device.ProcessDTC()

	case ruptela.COMMAND_EXTENDED_PROTOCOL_RECORDS:
		return devices.BEHAVE_REPLY_AND_CLOSE_CONN, device.ProcessExtended()

	case ruptela.COMMAND_GOBRAX_BLOCKED_IMEI:
		device.SetSuccess(true)
		return devices.BEHAVE_REPLY_AND_CLOSE_CONN, &customerrors.ErrBlockedImei

	}
	return devices.BEHAVE_REPLY_AND_CLOSE_CONN, &customerrors.ErrInvalidCommand
}
