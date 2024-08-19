package web

func (s *Server) ChangeState() {
	defer s.Close()
	switch s.state {
	case STATE_SET_TIMEOUT:
		s.timeoutState()
	case STATE_READDATA:
		s.readDataState()
	case STATE_HANDLEDATA:
		s.handleDataState()
	}
}
