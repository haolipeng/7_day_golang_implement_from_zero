package main

import (
	streamClientPb "7_day_golang_implement_from_zero/GeeGrpc/exercise/grpc/client_stream_rpc/pb"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type SimpleService struct {
}

//SayHello 实现接口
func (gs *SimpleService) SayHello(srv streamClientPb.StreamClient_SayHelloServer) error {
	//从stream流中获取消息
	for {
		res, err := srv.Recv()
		//判断流数据是否结束
		if err == io.EOF {
			//发送结果，并关闭
			fmt.Println("server received data finished")

			return srv.SendAndClose(&streamClientPb.SimpleReply{Message: "server done, close connection"})
		}
		if err != nil {
			return err
		}

		//打印流式结果
		log.Println(res.GetName())
	}

	return nil
}

func main() {
	//创建grpc server对象
	rpcServer := grpc.NewServer()

	//将GreeterServer服务注册到gRPC server
	streamClientPb.RegisterStreamClientServer(rpcServer, new(SimpleService))

	//创建 Listen，监听 TCP 的8081端口
	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	err = rpcServer.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
