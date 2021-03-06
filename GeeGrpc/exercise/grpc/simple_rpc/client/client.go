package main

import (
	simplePb "7_day_golang_implement_from_zero/GeeGrpc/exercise/grpc/simple_rpc/pb"
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
	defer conn.Close()

	//创建 Greeter 的客户端对象
	client := simplePb.NewGreeterClient(conn)

	//构造请求数据，并进行rpc调用
	req := simplePb.SimpleRequest{Name: "haolipeng"}

	reply, err := client.SayHello(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	//输出结果
	fmt.Println("message:", reply.Message)
}
