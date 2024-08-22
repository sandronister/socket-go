package web

import (
	"fmt"
	"log"
	"net"
	"time"
)

func (s *Server) Start() {
	address := s.getAddress()

	listener, err := net.ListenTCP("tcp", address)

	fmt.Println("Server started at", s.host+":"+s.port)

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		s.conn, err = listener.AcceptTCP()

		if err != nil {
			log.Fatal(err)
		}

		s.conn.SetDeadline(time.Now().Add(10 * time.Second))
		s.Handle()

	}
}
