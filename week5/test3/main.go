package main

import (
	"encoding/json"
	"fmt"
	"learngo/week5/test3/proto/pb"

	"google.golang.org/protobuf/proto"
)

type BookInfo struct {
	Name string `json:"name"`
}

func main() {
	req := pb.BookRequest{Name: "Go语言极简一本通"}
	b, err := proto.Marshal(&req)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
	fmt.Println(string(b))
	info := BookInfo{
		Name: "Go语言极简一本通",
	}
	jsonByte, err := json.Marshal(info)
	if err != nil {
		panic(err)
	}
	fmt.Println(jsonByte) // json 格式的二进制会更长
	fmt.Println(string(jsonByte))
}
