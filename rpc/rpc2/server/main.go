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
	*resp = "hello " + request
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}
	err = rpc.RegisterName("FoodService", new(FoodService))
	if err != nil {
		fmt.Println("register error:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			return
		}
		// rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}
}
