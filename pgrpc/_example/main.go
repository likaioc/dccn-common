package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/Ankr-network/dccn-common/pgrpc"
	"github.com/Ankr-network/dccn-common/pgrpc/_example/api"
	"google.golang.org/grpc"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	count := 10
	for idx := 0; idx < count; idx++ {
		go server()
	}
	client(count)
}

func server() {
	conn, err := pgrpc.NewServer("127.0.0.1:8899", 20*time.Second, nil)
	if err != nil {
		log.Fatalln("FAIL:", err)
	}

	log.Println("Server started")

	server := grpc.NewServer()
	api.RegisterPingServer(server, &api.Server{})
	if err := server.Serve(conn); err != nil {
		log.Fatalln("FAIL:", err)
	}
}

func client(serverCount int) {
	if err := pgrpc.InitClient("8899", 20*time.Second, log.New(os.Stderr, "", log.Ltime|log.Lshortfile)); err != nil {
		log.Fatalln(err)
	}

	log.Println("Client started")
	time.Sleep(time.Second) // a rough way to wait server connecting.
	clientKey := ""

	// Test: loop all server
	count := 0
	pgrpc.Each(func(key string, conn *grpc.ClientConn, err error) bool {
		if err != nil {
			log.Println(err)
			return true
		}

		out, err := api.NewPingClient(conn).SayHello(context.Background(),
			&api.PingMessage{Greeting: strconv.Itoa(count)})
		if err != nil {
			log.Fatalln("FAIL:", conn.Target(), err)
			return true
		}

		log.Println("SUCC:", conn.Target(), *out)
		clientKey = key // get a key, not care which picked
		count++
		return true
	})

	if count != serverCount {
		log.Fatalf("FAIL: %d servers fail.", serverCount-count)
	}

	// Test: parallel dial specified server
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			conn, err := pgrpc.Dial(clientKey)
			if err != nil {
				log.Fatalln("FAIL:", err)
			}
			_, err = api.NewPingClient(conn).SayHello(context.Background(),
				&api.PingMessage{Greeting: "Dial"})
			if err != nil {
				log.Fatalln("FAIL:", err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	log.Println("SUCC")
}
