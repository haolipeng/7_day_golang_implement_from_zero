package main

import (
	bothStreamPb "7_day_golang_implement_from_zero/GeeGrpc/exercise/grpc/bidirect_stream_rpc/pb"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"strconv"
)

type StreamService struct {
}

func (s *StreamService) Conversations(srv bothStreamPb.Stream_ConversationsServer) error {
	n := 1
	for {
		req, err := srv.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		err = srv.Send(&bothStreamPb.StreamReply{
			Message: "from stream server answer: the " + strconv.Itoa(n) + " question is " + req.Name,
		})
		if err != nil {
			return err
		}
		n++
		log.Printf("from stream client question: %s", req.Name)
	}
}

func main() {
	//创建grpc server对象
	rpcServer := grpc.NewServer()

	//将GreeterServer服务注册到gRPC server
	bothStreamPb.RegisterStreamServer(rpcServer, new(StreamService))

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
