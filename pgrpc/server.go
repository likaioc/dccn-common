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
	hook   Hook
}

func NewServer(addr string, hook Hook, conf *yamux.Config) (*Server, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, errors.Wrap(err, "tcp dail")
	}

	if hook != nil {
		if err := hook.OnAccept(addr, conn); err != nil {
			return nil, errors.Wrap(err, "on accept hook")
		}
	}

	if conf == nil {
		conf = yamux.DefaultConfig()
	}

	return &Server{
		addr:   addr,
		conn:   conn,
		config: conf,
		hook:   hook,
	}, nil
}

func (s *Server) Session() (net.Listener, error) {
	session, err := yamux.Server(s.conn, s.config)
	if err == nil {
		_, err = session.Ping()
	}
	if err != nil {
		s.conn.Close()
		s, err = NewServer(s.addr, s.hook, s.config)
		if err != nil {
			return nil, errors.Wrap(err, "session")
		}

		// retry
		session, err = yamux.Server(s.conn, s.config)
		if err != nil {
			return nil, errors.Wrap(err, "session")
		} else if _, err = session.Ping(); err != nil {
			return nil, errors.Wrap(err, "session")
		}
	}

	if s.hook != nil {
		if err := s.hook.OnBuild(s.addr, session); err != nil {
			return nil, errors.Wrap(err, "on build hook")
		}
	}

	go func() {
		<-session.CloseChan()
		if s.hook != nil {
			s.hook.OnClose(s.addr, session)
		}
	}()

	return session, nil
}
