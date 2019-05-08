package main

import (
	"log"
	"net"
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

type hook struct {
	pgrpc.EmptyHook
}

func (*hook) OnAccept(key *string, conn *net.Conn) error {
	(*conn).Write([]byte("PROXY TCP4 192.168.1.2 192.168.1.2 65535 65535\r\n"))
	return nil
}

func server(ip string) {
	conn, err := pgrpc.NewServer(ip+":8899", &hook{}, nil)
	if err != nil {
		log.Fatalln("FAIL:", err)
	}

	log.Println("Server started")

	server := grpc.NewServer()
	api.RegisterPingServer(server, &Server{})
	api.RegisterStreamPingServer(server, &StreamServer{})

	for {
		ln, err := conn.Session()
		if err != nil {
			log.Println("FAIL:", err)
			time.Sleep(2 * time.Second)
			continue
		}
		if err := server.Serve(ln); err != nil {
			log.Println("FAIL:", err)
		}
		time.Sleep(2 * time.Second)
	}
}
