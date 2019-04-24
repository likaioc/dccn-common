package main

import (
	"context"
	"log"
	"time"

	"github.com/Ankr-network/dccn-common/pgrpc/_example/api"
)

// Hello server
// Server represents the gRPC server
type Server struct {
}

// SayHello generates response to a Ping request
func (s *Server) SayHello(ctx context.Context, in *api.PingMessage) (*api.PingMessage, error) {
	log.Printf("Receive message: %s", in.Greeting)
	return &api.PingMessage{Greeting: "Hello " + in.Greeting}, nil
}

// Stream server
// Server represents the gRPC server
type StreamServer struct {
}

// SayHello generates response to a Ping request
func (*StreamServer) HelloStream(stream api.StreamPing_HelloStreamServer) error {
	go func() {
		for range time.Tick(2 * time.Second) {
			if err := stream.Send(&api.PingMessage{Greeting: "server side send"}); err != nil {
				log.Fatalln(err)
			}
			log.Println("Send", "server side send")

		}
	}()

	for {
		if msg, err := stream.Recv(); err != nil {
			log.Fatalln(err)
		} else {
			log.Println("Recv", msg.Greeting)
		}
	}
	return nil
}
