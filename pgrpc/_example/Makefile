IP:=
FLAG:=
MAKEFLAGS += --jobs=2

default: run_client run_server

.PHONY: server
server:
	GOOS=linux go build -o pgrpc_server $${PWD}/server
	docker ${FLAG} build -t test-pgrpc -f server/Dockerfile .

.PHONY: client
client:
	go build -o pgrpc_client $${PWD}/client

run_client: client
	./pgrpc_client

run_server: server
ifeq ("${IP}","")
	$(error IP is not set)
endif

	docker ${FLAG} run -it -e IP=$(IP) --rm test-pgrpc
