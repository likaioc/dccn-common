package pgrpc

import (
	"net"

	"github.com/hashicorp/yamux"
	"github.com/pkg/errors"
)

type Server struct {
	addr    string
	config  *yamux.Config
	session *yamux.Session
}

func NewServer(addr string, conf *yamux.Config) (*Server, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, errors.Wrap(err, "tcp dail")
	}

	if conf == nil {
		conf = yamux.DefaultConfig()
	}

	session, err := yamux.Server(conn, conf)
	if err != nil {
		return nil, errors.Wrap(err, "mux session")
	}

	return &Server{
		addr:    addr,
		config:  conf,
		session: session,
	}, nil
}

func (s *Server) Session() net.Listener {
	return s.session
}
