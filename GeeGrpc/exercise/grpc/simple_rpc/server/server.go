package main

import (
	simplePb "7_day_golang_implement_from_zero/GeeGrpc/exercise/grpc/simple_rpc/pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GreeterServer struct {
}

func (*GreeterServer) SayHello(ctx context.Context, req *simplePb.SimpleRequest) (*simplePb.SimpleReply, error) {
	//获取请求字段
	name := req.GetName()

	return &simplePb.SimpleReply{Message: "hello " + name}, nil
}

func main() {
	//创建grpc server对象
	rpcServer := grpc.NewServer()

	//将GreeterServer服务注册到gRPC server
	simplePb.RegisterGreeterServer(rpcServer, new(GreeterServer))

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
