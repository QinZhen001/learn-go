package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type FoodService struct {
}

func (f *FoodService) SayName(request string, resp *string) error {
	*resp = "您点的菜是" + request
	return nil

}

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = rpc.RegisterName("FoodService", &FoodService{})
	if err != nil {
		fmt.Println(err)
		return
	}

	for {
		// 等待请求，然后接受处理 ， wait request and conncetion， 返回下一个 请求者。
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		//rpc.ServeConn(conn)
		// json 格式 rpc服务
		//                NewServerCodec returns a new rpc.ServerCodec using JSON-RPC on conn.
		// ServeCodec is like ServeConn but uses the specified codec to
		// decode requests and encode responses.
		//使用goroutine并发
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
