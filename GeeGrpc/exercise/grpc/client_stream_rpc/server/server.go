package main

import (
	streamClientPb "7_day_golang_implement_from_zero/GeeGrpc/exercise/grpc/client_stream_rpc/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
)

type GreeterServer struct {
}

//SayHello 实现接口
func (gs *GreeterServer) SayHello(srv streamClientPb.Greeter_SayHelloServer) error {
	//从stream流中获取消息
	for true {
		res, err := srv.Recv()
		//判断流数据是否接触
		if err == io.EOF {
			return srv.SendAndClose(&streamClientPb.HelloReply{Message: "close connection"})
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
	streamClientPb.RegisterGreeterServer(rpcServer, new(GreeterServer))

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
