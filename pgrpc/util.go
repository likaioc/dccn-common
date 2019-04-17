package pgrpc

import (
	"net"

	"github.com/hashicorp/yamux"
)

type Hook interface {
	OnAccept(key string, conn net.Conn) error
	OnBuild(key string, conn *yamux.Session) error
	OnClose(key string, conn *yamux.Session)
}
