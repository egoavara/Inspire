package egn

import (
	"net"
	"net/http"
)

type Server struct {
	Address *net.TCPAddr
}

func NewServer() *Server {
	addr, _ := net.ResolveTCPAddr("", ":42424")
	return &Server{
		Address: addr,
	}
}
func (s *Server) Listen() (err error) {
	server := http.Server{
		Addr: s.Address.String(),
	}
	return server.ListenAndServe()
}