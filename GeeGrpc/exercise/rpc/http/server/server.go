package main

import (
	"7_day_golang_implement_from_zero/GeeGrpc/exercise/rpc/common"
	"log"
	"net/http"
	"net/rpc"
)

func main() {
	//注册rpc对象
	arith := new(common.Arith)
	rpc.Register(arith) //注册方法

	rpc.HandleHTTP() //注册http路由

	//在1234端口上启动http服务
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("server error:", err)
	}
}
