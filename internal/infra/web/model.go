package web

import (
	"net"

	"github.com/sandronister/socket-go/internal/dto"
	"github.com/sandronister/socket-go/internal/infra/handler"
)

type ProtocolStates uint8

const (
	STATE_SET_TIMEOUT ProtocolStates = 0x00
	STATE_READDATA    ProtocolStates = 0x01
	STATE_HANDLEDATA  ProtocolStates = 0x02
	STATE_SENDREPLY   ProtocolStates = 0x03
	STATE_CLOSE       ProtocolStates = 0x04
	STATE_MAX         ProtocolStates = STATE_CLOSE + 1
)

type Server struct {
	host      string
	port      string
	conn      *net.TCPConn
	state     ProtocolStates
	nextState ProtocolStates
	handler   handler.IHandler
	response  *dto.DeviceResponse
}
