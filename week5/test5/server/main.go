package main

import (
	"fmt"
	"learngo/week5/test5/proto/pb"
	"net"
	"sync"
	"time"

	"google.golang.org/grpc"
)

type FoodInfo struct {
}

// 服务端流模式
func (f *FoodInfo) SayName(request *pb.FoodStreamRequest, server pb.FoodService_SayNameServer) error {
	fmt.Println("SayName已请求")
	err := server.Send(&pb.FoodStreamResponse{
		Msg: "您点的菜是：" + request.Name,
	})
	if err != nil {
		return err
	}
	return nil
}

// 客户端流模式
func (f *FoodInfo) PostName(server pb.FoodService_PostNameServer) error {
	for {
		recv, err := server.Recv()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(recv.Name + ",您慢用！")
	}
	return nil
}
func (f *FoodInfo) FullStream(server pb.FoodService_FullStreamServer) error {
	//server.Send()
	//server.Recv()
	// 通过goroutine 处理阻塞
	// 通过 waitgroup 处理 等待
	var wg sync.WaitGroup
	wg.Add(2) //收 发
	c := make(chan string, 5)
	go func() {
		defer wg.Done()
		for {
			item, err := server.Recv()
			if err != nil {
				fmt.Println(err)
			}
			c <- item.Name
			fmt.Println("已下单：" + item.Name)
			time.Sleep(time.Second * 1)
		}
	}()

	go func() {
		defer wg.Done()
		for {
			foodName := <-c
			err := server.Send(&pb.FoodStreamResponse{Msg: "菜:" + foodName + "做好了"})
			if err != nil {
				fmt.Println(err.Error())
				// 是不是可以写一个recover 来返回报错。
				// 还是打日志 好一点 不会影响 协程
			}
			time.Sleep(time.Second * 1)
		}
	}()

	wg.Wait()
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":9091")
	if err != nil {
		panic(err)
	}
	server := grpc.NewServer()
	pb.RegisterFoodServiceServer(server, &FoodInfo{})
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}

}
