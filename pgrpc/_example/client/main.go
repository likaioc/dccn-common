package main

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/Ankr-network/dccn-common/pgrpc"
	"github.com/Ankr-network/dccn-common/pgrpc/_example/api"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	dial(":8899")

	once := sync.Once{}
	for range time.Tick(2 * time.Second) {
		keys := loopServers()
		log.Printf("connected servers(%d): %s", len(keys), keys)

		for _, key := range keys {
			dialServer(key)
		}
		once.Do(func() {
			// disable stream test
			// 	if len(keys) >= 1 {
			// 		streamMode(keys[0])
			// 	}
		})
	}
}

type hook struct {
	pgrpc.ProxyProtoHook
}

func (*hook) OnClose(key string, conn *pgrpc.Session) {
	log.Println(key, "is closed")
}

func dial(port string) {
	if err := pgrpc.InitClient(":8899", nil, new(hook)); err != nil {
		log.Fatalln(err)
	}

	log.Println("Client started")
}

func loopServers() (keys []string) {
	pgrpc.Each(func(key string, conn *grpc.ClientConn, err error) {
		if err != nil {
			log.Println(err)
			return
		}
		defer conn.Close()

		out, err := api.NewPingClient(conn).SayHello(context.Background(),
			&api.PingMessage{Greeting: key})
		if err != nil {
			log.Println("FAIL:", conn.Target(), err)
			return
		}

		pgrpc.Alias(key, "mock_name", true)

		log.Println("SUCC:", conn.Target(), *out)
		keys = append(keys, key)
	})

	return keys
}

func dialServer(key string) {
	conn, err := pgrpc.Dial(key)
	if err != nil {
		log.Println("FAIL:", err)
	}
	defer conn.Close()

	msg, err := api.NewPingClient(conn).SayHello(context.Background(),
		&api.PingMessage{Greeting: "Dial"})
	if err != nil {
		log.Println("FAIL:", err)
	}

	log.Println("SUCC:", *msg)
}

func streamMode(key string) {
	conn, err := pgrpc.Dial(key)
	if err != nil {
		log.Println("FAIL:", err)
	}
	defer conn.Close()

	stream, err := api.NewStreamPingClient(conn).HelloStream(context.Background())
	if err != nil {
		log.Println("FAIL:", err)
	}

	go func() {
		for range time.Tick(2 * time.Second) {
			if err := stream.Send(&api.PingMessage{Greeting: "client side send"}); err != nil {
				log.Println(err)
			}
			log.Println("Send", "client side send")
		}
	}()

	for {
		if msg, err := stream.Recv(); err != nil {
			log.Println(err)
		} else {
			log.Println("Recv", msg.Greeting)
		}
	}
}
