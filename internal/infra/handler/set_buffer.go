package handler

func (h *TcpHandler) SetBuffer(buffer []byte) {
	h.ReadBuffer = buffer
}
