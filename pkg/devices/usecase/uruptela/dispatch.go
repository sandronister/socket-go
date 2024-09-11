package uruptela

import (
	"github.com/sandronister/socket-go/pkg/devices"
	customerrors "github.com/sandronister/socket-go/pkg/devices/custom_errors"
)

func (u *UseRuptela) Dispatch(device devices.IDevice) *customerrors.Error {
	u.SendToBucket(device)
	msg, err := device.ToBytes()
	if err != nil {
		return &customerrors.ErrInvalidCommand
	}
	u.SendMessage(msg)
	return nil
}
