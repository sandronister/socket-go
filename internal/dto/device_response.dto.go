package dto

import "github.com/sandronister/socket-go/pkg/devices"

type DeviceResponse struct {
	Error          string
	Ack            []byte
	ProtocolBehave devices.ProtocolBehave
	Success        bool
	Imei           string
}
