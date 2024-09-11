package handler

import (
	"net/netip"
	"time"

	"github.com/sandronister/socket-go/pkg/devices/usecase"
)

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
	Write(b []byte) (int, error)
}

type IHandler interface {
	Handle(conn TCPAddrInterface, item []byte)
}

type TcpHandler struct {
	udevice    usecase.IDeviceUseCase
	Retries    int
	MaxRetries int
}
