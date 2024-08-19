package devices

import customerrors "github.com/sandronister/socket-go/pkg/devices/custom_errors"

type ProtocolBehave uint8

type HandleResponse struct {
	Imei           string
	Ack            []byte
	Success        bool
	Error          customerrors.Error
	ProtocolBehave ProtocolBehave
}

const (
	BEHAVE_UNKNOWN              ProtocolBehave = 0x00
	BEHAVE_CLOSE_CONN           ProtocolBehave = 0x00
	BEHAVE_REPLY_AND_KEEP_ALIVE ProtocolBehave = 0x01
	BEHAVE_REPLY_AND_CLOSE_CONN ProtocolBehave = 0x02
)
