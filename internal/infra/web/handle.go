package web

func (s *Server) Handle() {
	var channel chan []byte = make(chan []byte, 1)

	for range 4 {
		go s.LoadBuffer(channel)
	}

	s.Read(channel)

}
