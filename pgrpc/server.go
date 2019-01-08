package pgrpc

import (
	"log"
	"net"
	"time"

	"github.com/hashicorp/yamux"
	"github.com/pkg/errors"
)

type Server struct {
	addr    string
	config  *yamux.Config
	session *yamux.Session
}

func NewServer(addr string, KeepAliveInterval time.Duration, logger *log.Logger) (*Server, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, errors.Wrap(err, "tcp dail")
	}

	config := yamux.DefaultConfig()
	if KeepAliveInterval != 0 {
		config.KeepAliveInterval = KeepAliveInterval
	}
	if logger != nil {
		config.Logger = logger
		config.LogOutput = nil
	}

	session, err := yamux.Server(conn, config)
	if err != nil {
		return nil, errors.Wrap(err, "mux session")
	}

	return &Server{
		addr:    addr,
		config:  config,
		session: session,
	}, nil
}

func (s *Server) Session() net.Listener {
	return s.session
}
