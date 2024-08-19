package handler

import (
	"bufio"

	"github.com/sandronister/socket-go/internal/dto"
)

var i int = 1

func (h *TcpHandler) Handle(conn TCPAddrInterface) *dto.DeviceResponse {

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
