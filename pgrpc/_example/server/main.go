package main

import (
	"log"
	"time"

	"github.com/Ankr-network/dccn-common/pgrpc"
	"github.com/Ankr-network/dccn-common/pgrpc/_example/api"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	server()
}

func server() {
	conn, err := pgrpc.NewServer("127.0.0.1:8899", 20*time.Second, nil)
	if err != nil {
		log.Fatalln("FAIL:", err)
	}

	log.Println("Server started")

	server := grpc.NewServer()
	api.RegisterPingServer(server, &Server{})
	api.RegisterStreamPingServer(server, &StreamServer{})
	if err := server.Serve(conn); err != nil {
		log.Fatalln("FAIL:", err)
	}
}
