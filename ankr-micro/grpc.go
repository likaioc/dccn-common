package ankrmicro

import (
	"log"
	"net"

	//"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCService struct keep track of the server data struction
// of gRPC
type GRPCService struct {
	listener net.Listener
	s        *grpc.Server
}

// NewGRPCService creates a new gRPC service and returns the
// the gRPC Server
func NewGRPCService() GRPCService {
	s := GRPCService{}
	lis, err := net.Listen("tcp", config.Listen)
	s.listener = lis
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.s = grpc.NewServer()
	return s

}

// GetServer gets the underlying gPRC server instance
func (s *GRPCService) GetServer() *grpc.Server {
	return s.s
}

// Start will initiate the gPRC serving proces
func (s *GRPCService) Start() {
	// Register reflection service on gRPC server.
	reflection.Register(s.s)
	WriteLog("GRPCService Start @ " + config.Listen)
	if err := s.s.Serve(s.listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
