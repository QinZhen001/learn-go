package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	c, err := rpc.Dial("tcp", "localhost:9090")
	if err != nil {
		fmt.Println(err)
		return
	}

	reply := ""
	// Call invokes the named function, waits for it to complete, and returns its error status.
	err = c.Call("FoodService.SayName", "九转大肠", &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(1111)
	fmt.Println(reply)
}
