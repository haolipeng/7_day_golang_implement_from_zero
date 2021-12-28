package main

import (
	"7_day_golang_implement_from_zero/GeeGrpc/exercise/grpc/pb"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
)

func main() {
	//创建连接
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	//通过连接创建client对象
	client := pb.NewGreeterClient(conn)

	req := pb.HelloRequest{Name: "haolipeng"}

	//rpc调用
	reply, err := client.SayHello(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	//输出结果
	fmt.Println("message:", reply.Message)
}
