IP:=

default: server client run

.PHONY: server
server:
	GOOS=linux go build -o pgrpc_server $${PWD}/server
	docker build -t test-pgrpc -f server/Dockerfile .

.PHONY: client
client:
	go build -o pgrpc_client $${PWD}/client

run:
	ifeq ("${IP}","")
		$(error IP is not set)
	endif

	./pgrpc_client &
	docker run -it -e IP=$(IP) test-pgrpc
	pkill -9 pgrpc_client
