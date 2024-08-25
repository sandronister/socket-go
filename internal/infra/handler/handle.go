package handler

func (h *TcpHandler) Handle(conn TCPAddrInterface, item []byte) {
	response := h.udevice.Handle(item)
	conn.Write(response.Ack)
}
