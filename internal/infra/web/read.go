package web

import (
	"fmt"
	"time"
)

func (s *Server) Read(buffer chan<- []byte) {
	s.conn.SetReadDeadline(time.Now().Add(5 * time.Second))
	item := make([]byte, 1024)
	_, err := s.conn.Read(item)
	if err != nil {
		fmt.Println("Error reading data")
	}
	buffer <- item

}
