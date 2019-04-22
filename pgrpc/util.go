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

type EmptyHook struct{}

func (*EmptyHook) OnAccept(key string, conn net.Conn) error { return nil }
func (*EmptyHook) OnBuild(key string, conn *Session) error  { return nil }
func (*EmptyHook) OnClose(key string, conn *Session)        {}
