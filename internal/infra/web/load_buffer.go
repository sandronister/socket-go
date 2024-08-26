package web

import "fmt"

func (s *Server) LoadBuffer(buffer <-chan []byte) {
	for item := range buffer {
		s.getResponse(item)
		fmt.Println("Data handled")
		fmt.Println("----------------")
	}
}
