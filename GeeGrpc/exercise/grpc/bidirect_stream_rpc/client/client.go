package main

import (
	bothStreamPb "7_day_golang_implement_from_zero/GeeGrpc/exercise/grpc/bidirect_stream_rpc/pb"
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
	client := bothStreamPb.NewStreamClient(conn)

	stream, err := client.Conversations(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//向流中多次发送消息
	for n := 0; n < 5; n++ {
		err := stream.Send(&bothStreamPb.StreamRequest{Name: "stream client rpc " + strconv.Itoa(n)})
		if err != nil {
			log.Fatalf("stream request err: %v", err)
		}
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Conversations get stream err: %v", err)
		}
		// 打印返回值
		log.Println(res.Message)
	}

	//关闭流并获取返回的消息
	err = stream.CloseSend()
	if err != nil {
		log.Fatalf("Conversations close stream error:%v\n", err)
	}
}
