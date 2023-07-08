package main

import (
	"fmt"
	"net"
	"net/rpc"
)

type FoodService struct{}

func (f *FoodService) SayName(name string, resp *string) error {
	*resp = "您点的菜是" + name
	return nil
}

func main() {
	listener, err := net.Listen("tcp", ":9090")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = rpc.RegisterName("FoodService", &FoodService{})
	// Accept waits for and returns the next connection to the listener.
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println(err)
		return
	}
	// 服务链接，
	rpc.ServeConn(conn)
}
