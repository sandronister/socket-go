package handler

import "fmt"

func (h *TcpHandler) Handle(conn TCPAddrInterface, item []byte) {
	response := h.udevice.Handle(item)
	fmt.Println("Handling data", response)
	conn.Write(response.Ack)
}
