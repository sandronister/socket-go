package web

import (
	"bufio"
	"log"
	"net"
)

type Server struct {
	host string
	port string
	conn net.Conn
}

func NewServer(host, port string) *Server {
	return &Server{
		host: host,
		port: port,
	}
}

func (s *Server) Start() {
	listener, err := net.Listen("tcp", s.host+":"+s.port)

	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		s.conn, err = listener.Accept()

		if err != nil {
			log.Fatal(err)
		}

		go s.HandleConnection()

	}
}

func (s *Server) HandleConnection() {
	defer s.conn.Close()
	reader := bufio.NewReader(s.conn)

	for {
		msg, err := reader.ReadString('\n')

		if err != nil {
			log.Println(err)
			return
		}

		log.Println(msg)
	}

}
