package handler

import (
	"net/netip"

	"github.com/sandronister/go_broker/pkg/broker/types"
	"github.com/sandronister/socket-go/internal/usecase"
)

const bufferMaxSize = 1024

type AddrPortInterface interface {
	String() string
	Addr() netip.Addr
	Port() uint16
	WithZone(zone string) netip.AddrPort
}

type TCPAddrInterface interface {
	Read(p []byte) (n int, err error)
	Close() error
}

type IHandler interface {
	Handle(conn TCPAddrInterface)
	ReadTCP(conn TCPAddrInterface) (int, error)
}

type TcpHandler struct {
	broker     types.IBroker
	usecase    usecase.IUseCase
	ReadBuffer []byte
	Retries    int
	MaxRetries int
}
