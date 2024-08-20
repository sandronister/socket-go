package ruptela

import (
	"strings"

	customerrors "github.com/sandronister/socket-go/pkg/devices/custom_errors"
)

func (r *Device) IsValidImei() bool {
	return strings.HasPrefix(r.Header.Imei, string(MANUFACTURER_RUPTELA_IMEI_INIT))
}

func (r *Device) IsValidCommand() bool {
	recCommand := TCommandProtocol(r.Header.Command)

	for _, command := range allCommands {
		if command == recCommand {
			return true
		}
	}

	return false
}

func (r *Device) IsValidHeader() *customerrors.Error {
	if !r.IsValidImei() {
		return &customerrors.ErrInvalidImei
	}

	if !r.IsValidCommand() {
		return &customerrors.ErrInvalidCommand
	}

	return nil
}
