package main

import (
	"context"
	"fmt"
	"learngo/week6/test1/proto/pb"
	"net"

	"google.golang.org/grpc"
)

type DefaultValue struct {
}

func (d *DefaultValue) GetDefault(ctx context.Context, req *pb.DefaultRequest) (*pb.DefaultResponse, error) {
	fmt.Println(req)
	return &pb.DefaultResponse{
		Msg: "调用成功",
	}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterDefaultServiceServer(server, &DefaultValue{})
	listen, err := net.Listen("tcp", "0.0.0.0:9092")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}

}
