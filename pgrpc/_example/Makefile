IP:=
FLAG:=

default: server client run

.PHONY: server
server:
	GOOS=linux go build -o pgrpc_server $${PWD}/server
	docker ${FLAG} build -t test-pgrpc -f server/Dockerfile .

.PHONY: client
client:
	go build -o pgrpc_client $${PWD}/client

run:
ifeq ("${IP}","")
	$(error IP is not set)
endif

	./pgrpc_client &
	docker ${FLAG} run -it -e IP=$(IP) --rm test-pgrpc
	pkill -9 pgrpc_client
