package main

import (
	pb "7_day_golang_implement_from_zero/GeeGrpc/exercise/grpc/client_stream_rpc/pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"strconv"
)

func main() {
	//创建连接
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//创建 Greeter 的客户端对象
	client := pb.NewGreeterClient(conn)

	//调用服务端的SayHello方法，获得流
	stream, err := client.SayHello(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for n := 0; n < 5; n++ {
		stream.Send(&pb.HelloRequest{Name: "steam client rpc " + strconv.Itoa(n)})
	}

}
