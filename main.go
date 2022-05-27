package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"heartbeat-grpc/heartbeat_pb"
	"log"
	"net"
	"strconv"
)

type server struct {
	heartbeat_pb.UnimplementedHeartBeatServiceServer
}
type heartbeat struct {
	Bpm      int32
	Username string
}

func main() {
	opts := []grpc.ServerOption{}
	rpcServer := grpc.NewServer(opts...)
	heartbeat_pb.RegisterHeartBeatServiceServer(rpcServer, &server{})
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Fail to connet %v ", err)
	}

	if err := rpcServer.Serve(lis); err != nil {
		log.Fatalf("cannot serve the rpc %v", err)
	}

}
func (s *server) UserHeartBeat(ctx context.Context, in *heartbeat_pb.HeartBeatRequest) (*heartbeat_pb.HeartBeatResponse, error) {
	fmt.Println("User heart Beat received...!!!")
	bpm := in.GetHeartbeat().GetBpm()
	username := in.GetHeartbeat().GetUsername()
	result := "User: " + username + " HeartBeat BPM " + strconv.Itoa(int(bpm)) + "\n"
	return &heartbeat_pb.HeartBeatResponse{
		Result: result,
	}, nil
}


