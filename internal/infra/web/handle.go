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

func (s *Server) Prepare(buffer <-chan []byte) {
	for item := range buffer {
		s.handler.SetBuffer(item)
		s.handleDataState()
		fmt.Println("Data handled")
		fmt.Println("----------------")
	}
}

func (s *Server) Handle() {
	var channel chan []byte = make(chan []byte, 1)

	for range 10 {
		go s.Prepare(channel)
	}

	s.Read(channel)

}
