package pgrpc

import (
	"net"

	"github.com/hashicorp/yamux"
	"google.golang.org/grpc"
)

type Session struct {
	*yamux.Session

	Name string
	Opts []grpc.DialOption
}

type Hook interface {
	OnAccept(key string, conn net.Conn) error
	OnBuild(key string, conn *Session) error
	OnClose(key string, conn *Session)
}

type emptyHook struct{}

func (*emptyHook) OnAccept(key string, conn net.Conn) error { return nil }
func (*emptyHook) OnBuild(key string, conn *Session) error  { return nil }
func (*emptyHook) OnClose(key string, conn *Session)        {}
