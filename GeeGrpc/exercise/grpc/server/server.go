package main

import (
	"7_day_golang_implement_from_zero/GeeGrpc/exercise/grpc/pb"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GreeterServer struct {
}

func (*GreeterServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	//获取请求字段
	name := req.GetName()

	return &pb.HelloReply{Message: "hello " + name}, nil
}

func main() {
	rpcServer := grpc.NewServer()
	pb.RegisterGreeterServer(rpcServer, new(GreeterServer))

	l, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}

	err = rpcServer.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
