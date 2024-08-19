package handler

import (
	"bufio"
	"net"
)

var i int = 1

func (h *TcpHandler) Handle(conn *net.TCPConn) {

	reader := bufio.NewReader(conn)
	defer conn.Close()

	for {

		msg, err := reader.Read(h.ReadBuffer)

		if err != nil {
			return
		}

		if msg > 0 {
			h.SendMessage(h.ReadBuffer[:msg])
		}

	}
}
