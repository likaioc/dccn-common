.PHONY: proto build test run deploy local clean echo

GOPATH:=$(shell go env GOPATH)


APPS=$(wildcard */v1/*.proto)
RECURSIVE_PROTO = for app in $(APPS); do \
        @protoc -I. \
	    -I$(GOPATH)/src/github.com/Ankr-network/dccn-common/protos \
	    --micro_out=Mtask.proto=github.com/Ankr-network/dccn-common/protos/common:. \
	    --go_out=Mevent.proto=github.com/Ankr-network/dccn-common/protos/common:. \
	    $(subdir); \
        done


proto:
	@protoc -I. \
	--micro_out=. \
	--go_out=. \
	email/v1/micro/email.proto;

	@protoc -I. \
	--micro_out=. \
	--go_out=. \
	usermgr/v1/micro/usermgr.proto;

	@protoc -I. \
	-Icommon \
	--micro_out=Mcommon/common.proto=github.com/Ankr-network/dccn-common/protos/common,Mcommon/error.proto=github.com/Ankr-network/dccn-common/protos/common:. \
	--go_out=Mcommon/common.proto=github.com/Ankr-network/dccn-common/protos/common,Mcommon/error.proto=github.com/Ankr-network/dccn-common/protos/common:. \
	taskmgr/v1/micro/taskmgr.proto

	@protoc -I. \
	-Icommon \
	--micro_out=Mcommon/event.proto=github.com/Ankr-network/dccn-common/protos/common:. \
	--go_out=Mcommon/event.proto=github.com/Ankr-network/dccn-common/protos/common:. \
	dcmgr/v1/micro/dcmgr.proto

