package web

import (
	"fmt"
	"net"
)

func (s *Server) getAddress() *net.TCPAddr {
	connStr := fmt.Sprintf("%s:%s", s.host, s.port)
	address, _ := net.ResolveTCPAddr("tcp", connStr)
	return address
}
