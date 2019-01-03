// +build integration

package client_rpc

import (
	"testing"
)

const (
	port    = ":50051"
	address = "localhost:50051"
)

func TestConnect(t *testing.T) {
	lis, err := Connect(address)
	if lis == nil {
		t.Fail()
	}
	if err != nil {
		t.Fail()
	}
}

func TestCreateRequest(t *testing.T) {
	request := CreateRequest(1, "ed1605e17374bde6c68864d072c9f5c9", "us_west", "ca", "docker_image_name")
	if AddTask_util(address, request) != "Success" {
		t.Fail()
	}

}

func TestAddtask_util(t *testing.T) {
	request := CreateRequest(1, "ed1605e17374bde6c68864d072c9f5c9", "us_west", "ca", "docker_image_name")
	if AddTask_util(address, request) != "Success" {
		t.Fail()
	}
}

func TestTask_list_util(t *testing.T) {
	request := CreateRequest(2, "ed1605e17374bde6c68864d072c9f5c9")
	if Task_list_util(address, request) == nil {
		t.Fail()
	}

}

func TestCancelTask_util(t *testing.T) {
	request := CreateRequest(3, "ed1605e17374bde6c68864d072c9f5c9", "93")
	if CancelTask_util(address, request) != "Failure" {
		t.Fail()
	}
}
