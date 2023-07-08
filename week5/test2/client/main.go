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
		fmt.Println(err)
		return
	}
	reply := ""

	// XXX.MethodName
	// Json
	// NewClientWithCodec is like NewClient but uses the specified
	// codec to encode requests and decode responses.
	// NewClientCodec returns a new rpc.ClientCodec using JSON-RPC on conn.
	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(c))
	err = client.Call("FoodService.SayName", "爆炒小龙虾", &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(reply)

}
