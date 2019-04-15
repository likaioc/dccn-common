package pgrpc

import (
	"context"
	"net"
	"strings"
	"sync"

	"github.com/hashicorp/yamux"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Client struct {
	port   string
	config *yamux.Config
	conns  *sync.Map
	opts   []grpc.DialOption
}

var DefaultClient *Client

func InitClient(port string, conf *yamux.Config, defaultOpts ...grpc.DialOption) (err error) {
	DefaultClient, err = NewClient(port, conf, defaultOpts...)
	return
}
func NewClient(port string, conf *yamux.Config, defaultOpts ...grpc.DialOption) (*Client, error) {
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

			sess, err := yamux.Client(conn, conf)
			if err != nil {
				if conf.Logger != nil {
					conf.Logger.Println("mux conn", port, err)
				}
				continue
			}

			host, _, err := net.SplitHostPort(sess.RemoteAddr().String())
			if err != nil {
				if conf.Logger != nil {
					conf.Logger.Println("parse new connection address fail:", err)
				}
				continue
			}

			if conf.Logger != nil {
				conf.Logger.Println("new connection from:", host)
			}

			if conn, ok := conns.LoadOrStore(host, sess); ok {
				conns.Store(host, sess)
				go func() {
					<-sess.CloseChan()
					conns.Delete(host)
				}()

				if err := conn.(*yamux.Session).GoAway(); err != nil && conf.Logger != nil {
					conf.Logger.Printf("session(%s) go away fail: %s", host, err)
				}
			}
		}
	}()

	if len(defaultOpts) == 0 {
		defaultOpts = []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                conf.KeepAliveInterval,
				Timeout:             conf.ConnectionWriteTimeout,
				PermitWithoutStream: true,
			}),
		}
	}
	return &Client{
		port:   port,
		config: conf,
		conns:  conns,
		opts:   defaultOpts,
	}, nil
}

func Alias(key, alias string, force bool) error {
	return DefaultClient.Alias(key, alias, force)
}
func (c *Client) Alias(key, alias string, force bool) error {
	val, ok := c.conns.Load(key)
	if !ok {
		return errors.New("alias key not found")
	}

	if force {
		c.conns.Store(alias, val)
	} else if _, ok = c.conns.LoadOrStore(alias, val); ok {
		return errors.New("alias name has been occupied")
	}

	go func() {
		<-val.(*yamux.Session).CloseChan()
		c.conns.Delete(alias)
	}()
	return nil
}

func Dial(key string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return DefaultClient.Dial(key)
}
func (c *Client) Dial(key string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	conn, ok := c.conns.Load(key)
	if !ok {
		return nil, errors.New("no pgrpc connection target to " + key)
	}

	if len(opts) == 0 {
		opts = c.opts
	}
	opts = append(opts, grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return conn.(*yamux.Session).Open()
	}))

	return grpc.DialContext(context.Background(), key, opts...)
}

func Each(fn func(key string, conn *grpc.ClientConn, err error), opts ...grpc.DialOption) {
	DefaultClient.Each(fn, opts...)
}
func (c *Client) Each(fn func(key string, conn *grpc.ClientConn, err error), opts ...grpc.DialOption) {
	c.conns.Range(func(key, val interface{}) bool {

		if len(opts) == 0 {
			opts = c.opts
		}
		opts = append(opts, grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
			return val.(*yamux.Session).Open()
		}))
		conn, err := grpc.DialContext(context.Background(), key.(string), opts...)
		fn(key.(string), conn, errors.Wrapf(err, "pgrpc dial (%s)", key))

		return true
	})
}
