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
	if n > 0 {
		s.nextState = STATE_HANDLEDATA
	}
	return nil
}

func (s *Server) handleDataState() {
	s.handler.Handle(s.conn)
	s.nextState = STATE_CLOSE
}
