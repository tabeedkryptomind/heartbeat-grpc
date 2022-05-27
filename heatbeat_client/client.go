package main

import (
	"context"
	"google.golang.org/grpc"
	"heartbeat-grpc/heartbeat_pb"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("client fail to connect %v", err)
	}
	defer conn.Close()
	c := heartbeat_pb.NewHeartBeatServiceClient(conn)
	UserHeartBeat(c)
}

func UserHeartBeat(c heartbeat_pb.HeartBeatServiceClient) {
	hbReq := heartbeat_pb.HeartBeatRequest{
		Heartbeat: &heartbeat_pb.HeartBeat{
			Bpm:      75,
			Username: "Haseeb",
		},
	}
	c.UserHeartBeat(context.Background(), &hbReq)
}
