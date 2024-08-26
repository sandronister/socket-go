package web

func (s *Server) getResponse(item []byte) {
	defer s.conn.Close()
	s.handler.Handle(s.conn, item)
}
