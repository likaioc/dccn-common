package main

import (
	"log"
	"os"
	"time"

	"github.com/Ankr-network/dccn-common/pgrpc"
	"github.com/Ankr-network/dccn-common/pgrpc/_example/api"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	ip := os.Getenv("IP")
	if ip == "" {
		log.Fatalln("env IP must set")
	}
	server(ip)
}

func server(ip string) {
	conn, err := pgrpc.NewServer(ip+":8899", 20*time.Second, nil)
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
