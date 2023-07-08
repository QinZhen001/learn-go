package main

import (
	"context"
	"learngo/week5/test4/proto/pb"
	"net"

	"google.golang.org/grpc"
)

type BookInfo struct {
}

func (b *BookInfo) Study(ctx context.Context, req *pb.BookRequest) (*pb.BookResponse, error) {
	return &pb.BookResponse{Msg: "我要学习" + req.Name}, nil
}

func main() {
	server := grpc.NewServer()
	pb.RegisterStudyServer(server, &BookInfo{})
	linten, err := net.Listen("tcp", "0.0.0.0:9090")
	if err != nil {
		panic(err)
	}
	err = server.Serve(linten)
	if err != nil {
		panic(err)
	}
}
