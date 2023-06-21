package main

import (
	"context"
	pb "learngo/rpc/rpc3/proto"

	"google.golang.org/grpc"
)

type FoodInfo struct {
}

func SayName(ctx context.Context, in *pb.FoodStreamRequest, opts ...grpc.CallOption) (pb.FoodService_SayNameClient, error)

// 客户端流模式
func PostName(ctx context.Context, opts ...grpc.CallOption) (pb.FoodService_PostNameClient, error)

// 双向流模式
func FullName(ctx context.Context, opts ...grpc.CallOption) (pb.FoodService_FullNameClient, error)
