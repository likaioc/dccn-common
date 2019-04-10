package pgrpc

import (
	"net"

	"github.com/hashicorp/yamux"
	"github.com/pkg/errors"
)

type Server struct {
	addr   string
	conn   net.Conn
	config *yamux.Config
}

func NewServer(addr string, conf *yamux.Config) (*Server, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, errors.Wrap(err, "tcp dail")
	}

	if conf == nil {
		conf = yamux.DefaultConfig()
	}

	return &Server{
		addr:   addr,
		conn:   conn,
		config: conf,
	}, nil
}

func (s *Server) Session() (net.Listener, error) {
	session, err := yamux.Server(s.conn, s.config)
	if err != nil {
		s.conn.Close()
		s, err = NewServer(s.addr, s.config)
		if err != nil {
			return nil, errors.Wrap(err, "session")
		}

		// retry
		session, err = yamux.Server(s.conn, s.config)
		if err != nil {
			return nil, errors.Wrap(err, "session")
		}
	}

	return session, nil
}
