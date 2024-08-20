package uruptela

import (
	"github.com/sandronister/socket-go/pkg/devices"
	customerrors "github.com/sandronister/socket-go/pkg/devices/custom_errors"
	"github.com/sandronister/socket-go/pkg/devices/entities/ruptela"
)

func (u *usecase_t) getParameter(device devices.IDevice) (customerrors.Error, devices.ProtocolBehave) {
	cmd := device.ProcessHeader()

	switch cmd {
	case ruptela.COMMAND_DYNAMIC_IDENTFIC_PACKET:
		return nil, devices.BEHAVE_REPLY_AND_KEEP_ALIVE
	}
}
