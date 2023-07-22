package main

import (
	"context"
	"fmt"
	"learngo/week6/test1/proto/pb"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:9092", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	client := pb.NewDefaultServiceClient(conn)
	res, err := client.GetDefault(context.Background(), &pb.DefaultRequest{})
	if err != nil {
		panic(err)
	}
	fmt.Println(1111, res.Msg)
}
