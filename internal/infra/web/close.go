package web

func (s *Server) Close() {
	if s.conn != nil {
		s.conn.Close()
	}
}
