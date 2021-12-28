package main

import (
	streamServerPb "7_day_golang_implement_from_zero/GeeGrpc/exercise/grpc/server_stream_rpc/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
)

type GreeterServer struct {
}

//SayHello 实现接口
func (gs *GreeterServer) SayHello(req *streamServerPb.HelloRequest, srv streamServerPb.Greeter_SayHelloServer) error {
	var err error

	// 向流中发送消息， 默认每次send送消息最大长度为`math.MaxInt32`bytes
	// 构造不同数据，发送多个响应给客户端
	for n := 0; n < 5; n++ {
		reply := &streamServerPb.HelloReply{Message: req.GetName() + strconv.Itoa(n)}

		err = srv.Send(reply)
		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	//创建grpc server对象
	rpcServer := grpc.NewServer()

	//将GreeterServer服务注册到gRPC server
	streamServerPb.RegisterGreeterServer(rpcServer, new(GreeterServer))

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
