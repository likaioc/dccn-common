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
	hook   Hook
}

var DefaultClient *Client

func InitClient(port string, conf *yamux.Config, hook Hook, opts ...grpc.DialOption) (err error) {
	DefaultClient, err = NewClient(port, conf, hook, opts...)
	return
}
func NewClient(port string, conf *yamux.Config, hook Hook, opts ...grpc.DialOption) (*Client, error) {
	// init yamux
	if conf == nil {
		conf = yamux.DefaultConfig()
	}
	if hook == nil {
		hook = new(EmptyHook)
	}
	if len(opts) == 0 {
		opts = []grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                conf.KeepAliveInterval,
				Timeout:             conf.ConnectionWriteTimeout,
				PermitWithoutStream: true,
			}),
		}
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

			host, _, err := net.SplitHostPort(conn.RemoteAddr().String())
			if err != nil {
				if conf.Logger != nil {
					conf.Logger.Println("parse new connection address fail:", err)
				}
				continue
			}

			if hook != nil {
				if err := hook.OnAccept(host, conn); err != nil {
					if conf.Logger != nil {
						conf.Logger.Println("on accept hook:", err)
					}
					continue
				}
			}

			session, err := yamux.Client(conn, conf)
			if err != nil {
				if conf.Logger != nil {
					conf.Logger.Println("mux conn", port, err)
				}
				continue
			}

			sess := &Session{
				Session: session,
				Name:    host,
				Opts:    opts,
			}
			if err := hook.OnBuild(host, sess); err != nil {
				if conf.Logger != nil {
					conf.Logger.Println("on build hook:", err)
				}
				continue
			}

			if conf.Logger != nil {
				conf.Logger.Println("new connection from:", host)
			}

			if session, ok := conns.LoadOrStore(host, sess); ok {
				conns.Store(host, sess) // force store

				go func() {
					<-sess.CloseChan()
					if val, ok := conns.Load(host); ok && val.(*Session).Name == host {
						hook.OnClose(host, val.(*Session))
					}
					conns.Delete(host)
				}()

				if err := session.(*Session).GoAway(); err != nil && conf.Logger != nil {
					conf.Logger.Printf("session(%s) go away fail: %s", host, err)
				}
			}
		}
	}()

	return &Client{
		port:   port,
		config: conf,
		conns:  conns,
		hook:   hook,
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

	if _, ok = c.conns.LoadOrStore(alias, val); !ok {
		val.(*Session).Name = alias // concurrent unsafe
	} else if force {
		c.conns.Store(alias, val)
		val.(*Session).Name = alias // concurrent unsafe
		return nil                  // avoid rerun hook
	} else {
		return errors.New("alias name has been occupied")
	}

	go func() {
		<-val.(*Session).CloseChan()
		if val, ok := c.conns.Load(alias); ok && val.(*Session).Name == alias {
			c.hook.OnClose(alias, val.(*Session))
		}
		c.conns.Delete(alias)
	}()
	return nil
}

func Dial(key string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	return DefaultClient.Dial(key)
}
func (c *Client) Dial(key string, opts ...grpc.DialOption) (*grpc.ClientConn, error) {
	sess, ok := c.conns.Load(key)
	if !ok {
		return nil, errors.New("no pgrpc connection target to " + key)
	}

	if len(opts) == 0 {
		opts = sess.(*Session).Opts
	}
	opts = append(opts, grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
		return sess.(*Session).Open()
	}))

	return grpc.DialContext(context.Background(), key, opts...)
}

func Each(fn func(key string, conn *grpc.ClientConn, err error), opts ...grpc.DialOption) {
	DefaultClient.Each(fn, opts...)
}
func (c *Client) Each(fn func(key string, conn *grpc.ClientConn, err error), opts ...grpc.DialOption) {
	wg := sync.WaitGroup{}
	defer wg.Wait()
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	c.conns.Range(func(key, val interface{}) bool {
		if v, ok := c.conns.Load(key); !ok || v.(*Session).Name != key.(string) {
			return true
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			if len(opts) == 0 {
				opts = val.(*Session).Opts
			}
			opts = append(opts, grpc.WithContextDialer(func(context.Context, string) (net.Conn, error) {
				sess := val.(*Session)
				if _, err := sess.Ping(); err != nil {
					return nil, err
				}
				return sess.Open()
			}))

			conn, err := grpc.DialContext(ctx, key.(string), opts...)
			fn(key.(string), conn, errors.Wrapf(err, "pgrpc dial (%s)", key))
		}()

		return true
	})
}
