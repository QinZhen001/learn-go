package main

import (
	"fmt"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	c, err := net.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	reply := ""

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c))
	err = client.Call("FoodService.SayName", "tom", &reply)
	if err != nil {
		fmt.Println("call error:", err)
		return
	}
	fmt.Println(reply)
}
