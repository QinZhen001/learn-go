package main

import (
	"context"
	"fmt"
	"learngo/week5/test4/proto/pb"

	"google.golang.org/grpc"
)

func main() {
	// 客户端知道对方的地址拨号进去，返回一个链接 conn
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	// 通过链接创建我的客户端
	client := pb.NewStudyClient(conn)
	// 客户端调用服务端的方法，返回结果
	res, err := client.Study(context.Background(), &pb.BookRequest{Name: "从0到Go语言微服务架构师"})
	if err != nil {
		panic(err)
	}
	fmt.Println(res.Msg)
}
