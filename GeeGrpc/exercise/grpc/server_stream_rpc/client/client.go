package main

import (
	streamServerPb "7_day_golang_implement_from_zero/GeeGrpc/exercise/grpc/server_stream_rpc/pb"
	"context"
	"google.golang.org/grpc"
	"io"
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
	client := streamServerPb.NewStreamServerClient(conn)

	//构造请求数据
	req := streamServerPb.SimpleRequest{
		Name: "stream server grpc ",
	}

	//调用服务端的SayHello方法，获得流
	stream, err := client.SayHello(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
	}

	for {
		//Recv() 方法接收服务端消息，默认每次Recv()最大消息长度为`1024*1024*4`bytes(4M)
		res, err := stream.Recv()
		// 判断消息流是否已经结束
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("ListStr get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.Message)
		// break
	}
	// //可以使用CloseSend()关闭stream，这样服务端就不会继续产生流消息
	// //调用CloseSend()后，若继续调用Recv()，会重新激活stream，接着之前结果获取消息
	// stream.CloseSend()

}
