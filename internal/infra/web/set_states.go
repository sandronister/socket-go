package web

import (
	"log"
	"time"
)

func (s *Server) timeoutState() {
	err := s.conn.SetDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		log.Println("Error setting deadline")
	}
	s.nextState = STATE_READDATA
}

func (s *Server) readDataState() error {
	n, err := s.handler.ReadTCP(s.conn)
	if err != nil {
		s.nextState = STATE_CLOSE
		return err
	}
	if n != nil {
		s.nextState = STATE_HANDLEDATA
	}
	return nil
}

func (s *Server) handleDataState() {
	s.response = s.handler.Handle()
	s.nextState = STATE_SENDREPLY
}

func (s *Server) sendReplyState() {
	s.nextState = STATE_CLOSE
	n, err := s.conn.Write(s.response.Ack)

	if err == nil || n > 0 {
		s.nextState = STATE_READDATA
	}

}
