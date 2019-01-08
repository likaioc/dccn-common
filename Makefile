# To use make toolchain, you must install protoc and protobuf manually.

default: gen-k8s gen-cli

gen-k8s:
	protoc --go_out=plugins=grpc:. protocol/k8s/*.proto

gen-cli:
	protoc --go_out=plugins=grpc:. protocol/cli/*.proto