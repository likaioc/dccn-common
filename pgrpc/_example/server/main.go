package main

import (
	"log"
	"os"

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
	conn, err := pgrpc.NewServer(ip+":8899", nil, nil)
	if err != nil {
		log.Fatalln("FAIL:", err)
	}

	log.Println("Server started")

	server := grpc.NewServer()
	api.RegisterPingServer(server, &Server{})
	api.RegisterStreamPingServer(server, &StreamServer{})
	ln, err := conn.Session()
	if err != nil {
		log.Fatalln("FAIL:", err)
	}
	if err := server.Serve(ln); err != nil {
		log.Fatalln("FAIL:", err)
	}
}
