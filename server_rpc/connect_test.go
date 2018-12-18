
package server_rpc

import (
	"testing"
)


const port = ":50051"
const address  = "localhost:50051"


func TestConnect(t *testing.T){
	lis, s := Connect(address)
	if lis == nil{
		t.Fail()
	}
	if s == nil {
        t.Fail()
    }

}
