package pgrpc

import (
	"context"
	"log"
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
	port        string
	config      *yamux.Config
	grpcOptions []grpc.DialOption
	conns       *sync.Map
}

var DefaultClient *Client

func InitClient(port string, KeepAliveInterval time.Duration, logger *log.Logger) (err error) {
	DefaultClient, err = NewClient(port, KeepAliveInterval, logger)
	return
}

func NewClient(port string, KeepAliveInterval time.Duration, logger *log.Logger) (*Client, error) {
	// init yamux
	config := yamux.DefaultConfig()
	if KeepAliveInterval != 0 {
		config.KeepAliveInterval = KeepAliveInterval
	}
	if logger != nil {
		config.Logger = logger
		config.LogOutput = nil
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
				if logger != nil {
					logger.Println("accept", port, err)
				}
				continue
			}

			muxConn, err := yamux.Client(conn, config)
			if err != nil {
				if logger != nil {
					logger.Println("mux conn", port, err)
				}
				continue
			}

			if logger != nil {
				logger.Println("new connection from:", muxConn.RemoteAddr().String())
			}

			if conn, ok := conns.LoadOrStore(muxConn.RemoteAddr().String(), muxConn); ok {
				conns.Store(muxConn.RemoteAddr().String(), muxConn)

				if err := conn.(*yamux.Session).GoAway(); err != nil && logger != nil {
					logger.Printf("session(%s) go away fail: %s", muxConn.RemoteAddr().String(), err)
				}
			}
		}
	}()

	return &Client{
		port:   port,
		config: config,
		grpcOptions: []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                config.KeepAliveInterval,
				Timeout:             config.ConnectionWriteTimeout,
				PermitWithoutStream: true,
			}),
		},
		conns: conns,
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

	grpcOptions := append(c.grpcOptions,
		grpc.WithDialer(func(string, time.Duration) (net.Conn, error) {
			return conn.(*yamux.Session).Open()
		}))
	return grpc.DialContext(context.Background(), key, grpcOptions...)
}

func Each(fn func(key string, conn *grpc.ClientConn, err error) (stop bool)) {
	DefaultClient.Each(fn)
}

func (c *Client) Each(fn func(key string, conn *grpc.ClientConn, err error) bool) {
	c.conns.Range(func(key, val interface{}) bool {
		grpcOptions := append(c.grpcOptions,
			grpc.WithDialer(func(string, time.Duration) (net.Conn, error) {
				return val.(*yamux.Session).Open()
			}))
		conn, err := grpc.DialContext(context.Background(), key.(string), grpcOptions...)

		return fn(key.(string), conn, errors.Wrapf(err, "pgrpc dial (%s)", key))
	})
}
