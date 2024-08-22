package web

import (
	"time"
)

func (s *Server) ChangeState() {
	s.nextState = STATE_SET_TIMEOUT
	s.state = STATE_SET_TIMEOUT
	defer s.Close()
	for {
		switch s.state {
		case STATE_SET_TIMEOUT:
			s.timeoutState()
		case STATE_READDATA:
			s.readDataState()
		case STATE_HANDLEDATA:
			s.handleDataState()
		case STATE_SENDREPLY:
			s.sendReplyState()
		case STATE_CLOSE:
			time.Sleep(time.Microsecond * 10)
			s.handler.ClearBuffer()
			return
		}
		s.state = s.nextState
	}

}
