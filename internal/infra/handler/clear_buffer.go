package handler

func (s *TcpHandler) ClearBuffer() {
	s.ReadBuffer = make([]byte, bufferMaxSize)
}
