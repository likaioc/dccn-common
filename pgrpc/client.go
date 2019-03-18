package pgrpc

import (
	"context"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/hashicorp/yamux"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Client struct {
	port   string
	config *yamux.Config
	conns  *sync.Map
}

var DefaultClient *Client

func InitClient(port string, conf *yamux.Config) (err error) {
	DefaultClient, err = NewClient(port, conf)
	return
}

func NewClient(port string, conf *yamux.Config) (*Client, error) {
	// init yamux
	if conf == nil {
		conf = yamux.DefaultConfig()
	}

	// dial
	if !strings.Contains(port, ":") {
		port = ":" + port
	}

	ln, err := net.Listen("tcp", port)
	if err != nil {
		return nil, errors.Wrap(err, "pgrpc listen")
	}

	conns := &sync.Map{}
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				if conf.Logger != nil {
					conf.Logger.Println("accept", port, err)
				}
				continue
			}

			muxConn, err := yamux.Client(conn, conf)
			if err != nil {
				if conf.Logger != nil {
					conf.Logger.Println("mux conn", port, err)
				}
				continue
			}

			if conf.Logger != nil {
				conf.Logger.Println("new connection from:", muxConn.RemoteAddr().String())
			}

			if conn, ok := conns.LoadOrStore(muxConn.RemoteAddr().String(), muxConn); ok {
				conns.Store(muxConn.RemoteAddr().String(), muxConn)

				if err := conn.(*yamux.Session).GoAway(); err != nil && conf.Logger != nil {
					conf.Logger.Printf("session(%s) go away fail: %s", muxConn.RemoteAddr().String(), err)
				}
			}
		}
	}()

	return &Client{
		port:   port,
		config: conf,
		conns:  conns,
	}, nil
}

func Dial(key string) (*grpc.ClientConn, error) {
	return DefaultClient.Dial(key)
}

func (c *Client) Dial(key string) (*grpc.ClientConn, error) {
	conn, ok := c.conns.Load(key)
	if !ok {
		return nil, errors.New("no pgrpc connection target to " + key)
	}

	return grpc.DialContext(context.Background(), key, grpc.WithInsecure(),
		grpc.WithDialer(func(string, time.Duration) (net.Conn, error) {
			return conn.(*yamux.Session).Open()
		}))
}

func Each(fn func(key string, conn *grpc.ClientConn, err error) (stop bool)) {
	DefaultClient.Each(fn)
}

func (c *Client) Each(fn func(key string, conn *grpc.ClientConn, err error) bool) {
	c.conns.Range(func(key, val interface{}) bool {
		conn, err := grpc.DialContext(context.Background(), key.(string), grpc.WithInsecure(),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                c.config.KeepAliveInterval,
				Timeout:             c.config.ConnectionWriteTimeout,
				PermitWithoutStream: true,
			}),
			grpc.WithDialer(func(string, time.Duration) (net.Conn, error) {
				return val.(*yamux.Session).Open()
			}))

		return fn(key.(string), conn, errors.Wrapf(err, "pgrpc dial (%s)", key))
	})
}
