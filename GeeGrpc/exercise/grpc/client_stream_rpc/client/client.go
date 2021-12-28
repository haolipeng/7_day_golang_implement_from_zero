package main

import (
	clientStreamPb "7_day_golang_implement_from_zero/GeeGrpc/exercise/grpc/client_stream_rpc/pb"
	"context"
	"google.golang.org/grpc"
	"io"
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
	client := clientStreamPb.NewStreamClientClient(conn)

	stream, err := client.SayHello(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//向流中多次发送消息
	for i := 0; i < 5; i++ {
		req := &clientStreamPb.StreamRequest{
			Name: "stream client rpc " + strconv.Itoa(i),
		}

		err = stream.Send(req)
		//发送也要检测EOF，当服务端在消息没接收完前主动调用SendAndClose()关闭stream
		//此时客户端还执行Send()，则会返回EOF错误，所以这里需要加上io.EOF判断
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
	}
	//关闭流并获取返回的消息
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("get response err: %v", err)
	}
	log.Println(res)
}
