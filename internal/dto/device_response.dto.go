package dto

import (
	"github.com/sandronister/socket-go/pkg/devices"
	customerrors "github.com/sandronister/socket-go/pkg/devices/custom_errors"
)

type DeviceResponse struct {
	Error          *customerrors.Error
	Ack            []byte
	ProtocolBehave devices.ProtocolBehave
	Success        bool
	Imei           string
}
