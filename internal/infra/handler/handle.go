package handler

import (
	"bufio"

	"github.com/sandronister/socket-go/pkg/devices"
)

var i int = 1

func (h *TcpHandler) Handle(conn TCPAddrInterface) *devices.HandleResponse {

	reader := bufio.NewReader(conn)
	defer conn.Close()

	for {

		msg, err := reader.Read(h.ReadBuffer)

		if err != nil {
			return nil
		}

		if msg > 0 {
			h.SendMessage(h.ReadBuffer[:msg])
		}

	}
}
