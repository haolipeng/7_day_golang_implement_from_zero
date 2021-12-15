package main

import (
	"7_day_golang_implement_from_zero/GeeGrpc/exercise/rpc/common"
	"log"
	"net"
	"net/rpc"
)

func main() {
	//1.生成监听器
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("net.Listen error:", err)
	}

	//2.创建对象，并注册方法到rpc中
	arith := new(common.Arith)
	rpc.Register(arith) //注册方法

	//3.在监听器上监听连接
	rpc.Accept(l)
}
