package web

import (
	"fmt"
	"time"
)

func (s *Server) ChangeState() {
	defer s.Close()
	for {
		fmt.Println("Current state:", s.state)
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
			return
		}
		s.state = s.nextState
	}

}
