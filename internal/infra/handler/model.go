package handler

import (
	"net/netip"
	"time"

	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/socket-go/internal/dto"
	"github.com/sandronister/socket-go/pkg/devices/usecase"
)

const bufferMaxSize = 1024

type AddrPortInterface interface {
	String() string
	Addr() netip.Addr
	Port() uint16
	WithZone(zone string) netip.AddrPort
}

type TCPAddrInterface interface {
	Read(b []byte) (n int, err error)
	Close() error
	SetReadDeadline(t time.Time) error
}

type IHandler interface {
	Handle() *dto.DeviceResponse
	ReadTCP(conn TCPAddrInterface) (int, error)
	ClearBuffer()
}

type TcpHandler struct {
	broker     types.IBroker
	udevice    usecase.IDeviceUseCase
	ReadBuffer []byte
	Retries    int
	MaxRetries int
}
