package main

import (
	"context"
	"fmt"
	"learngo/week6/test2/proto/pb"
	"net"

	"google.golang.org/grpc"
)

type Mistake struct {
}

func (m *Mistake) MakeMistake(ctx context.Context, req *pb.MistakeRequest) (*pb.MistakeResponse, error) {
	s := fmt.Sprintf("title:%s, subTitle:%s", req.Title, req.SubTitle)
	fmt.Println(s)
	return &pb.MistakeResponse{}, nil

}

func main() {
	server := grpc.NewServer()
	pb.RegisterMistakeServiceServer(server, &Mistake{})
	listen, err := net.Listen("tcp", "0.0.0.0:9093")
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
