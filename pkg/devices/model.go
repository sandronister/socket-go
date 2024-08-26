package devices

import (
	customerrors "github.com/sandronister/socket-go/pkg/devices/custom_errors"
	"github.com/sandronister/socket-go/pkg/devices/entities/ruptela"
	"github.com/sandronister/socket-go/pkg/devices/service"
)

type ProtocolBehave uint8

const (
	BEHAVE_UNKNOWN              ProtocolBehave = 0x00
	BEHAVE_CLOSE_CONN           ProtocolBehave = 0x00
	BEHAVE_REPLY_AND_KEEP_ALIVE ProtocolBehave = 0x01
	BEHAVE_REPLY_AND_CLOSE_CONN ProtocolBehave = 0x02
)

type IDevice interface {
	IsValidHeader() *customerrors.Error
	IsValidImei() bool
	IsValidCommand() bool
	ProcessHeader(blacklist service.IBlackListService) ruptela.TCommandProtocol
	ProcessDynamic() *customerrors.Error
	ProcessDTC() *customerrors.Error
	ProcessExtended() *customerrors.Error
	SetSuccess(bool)
	GetImei() string
	GetBuffer() []byte
	ToBytes() ([]byte, error)
}
