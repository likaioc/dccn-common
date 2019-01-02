package client_rpc

import (
		"strconv"
		"log"
		"fmt"
		"time"
	    "google.golang.org/grpc"
	    "golang.org/x/net/context"
	    pb "github.com/Ankr-network/dccn-common/protocol/cli"
)

func Connect(address_1 string) (*grpc.ClientConn, error) {
	return grpc.Dial(address_1, grpc.WithInsecure())
}
func AddTaskRequest_util(Usertoken1 string, Name1 string, Type1 string, Datacenter1 string) pb.AddTaskRequest {
	datacenter, err := strconv.ParseInt(Datacenter1, 10, 64)
	if err != nil {
		log.Fatalf("Client: cannot cover %v", err)
		}
	return pb.AddTaskRequest{Name: Name1, Type: Type1, Datacenterid: int64(datacenter), Usertoken: Usertoken1}
}

func TaskListRequest_util(Usertoken1 string) pb.TaskListRequest {
	return pb.TaskListRequest{Usertoken: Usertoken1}
}

func CancelTaskRequest_util(Usertoken1 string, Taskid1 int64) pb.CancelTaskRequest {
	return pb.CancelTaskRequest{Usertoken: Usertoken1, Taskid: Taskid1}
}

func CreateRequest(Whichrequest string, Message... string)(interface{}){
	fmt.Printf("received request : %d \n", Whichrequest)
	switch Whichrequest {
	case "AddTaskRequest":
		return AddTaskRequest_util(Message[0], Message[1], Message[2], Message[3])
//	}
//	if(Whichrequest == 2)
//	{
	case "TaskListRequest":
		return TaskListRequest_util(Message[0])

//	}
	//if(Whichrequest == 3){
	case "CancelTaskRequest":
		Taskid, err := strconv.ParseInt(Message[1], 10, 64)
		if err != nil {
			return nil
		}
		return CancelTaskRequest_util(Message[0], Taskid)
	}
	return nil
}

func AddTask_util(address string, request interface{}) string {
	var request1 pb.AddTaskRequest
	request1 = request.(pb.AddTaskRequest)
	conn, err := Connect(address)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDccncliClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	r, err := c.AddTask(ctx, &request1)
	if err != nil {
		log.Fatalf("Client: could not send: %v", err)
	}

	fmt.Printf("received Status : %s \n", r.Status)

	return r.Status

}

func Task_list_util(address string, request interface{}) []*pb.TaskInfo {
	var request1 pb.TaskListRequest
	request1 = request.(pb.TaskListRequest)
	conn, err := Connect(address)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDccncliClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	r, err := c.TaskList(ctx, &request1)
	if err != nil {
		log.Fatalf("Client: could not send: %v", err)
	}

	return r.Tasksinfo
	// todo when have new proto
}

func CancelTask_util(address string, request interface{}) string {
	var request1 pb.CancelTaskRequest
	request1 = request.(pb.CancelTaskRequest)
	conn, err := Connect(address)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewDccncliClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	r, err := c.CancelTask(ctx, &request1)
	if err != nil {
		log.Fatalf("Client: could not send: %v", err)
	}

	fmt.Printf("received Status : %s \n", r.Status)

	return r.Status

}
