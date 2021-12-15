package main

import (
	"7_day_golang_implement_from_zero/GeeGrpc/exercise/rpc/common"
	"fmt"
	"log"
	"net/rpc"
	"time"
)

func main() {
	//1.连接到服务器的监听地址
	client, err := rpc.DialHTTP("tcp", ":1234")
	if err != nil {
		log.Fatal("rpc.DialHTTP error:", err)
	}

	var quo common.Quetient
	args := common.Args{
		A: 20,
		B: 4,
	}
	//2.准备请求参数和返回参数
	DivDoneCall := client.Go("Arith.Divide", &args, &quo, nil)

	//3.创建定时器，等待10s
	t := time.NewTimer(10 * time.Second)

	for {
		select {
		case done := <-DivDoneCall.Done:
			if done.Error != nil {
				fmt.Println("Multiply error:", done.Error)
			} else {
				fmt.Printf("Divide: %d/%d=%d...%d\n", args.A, args.B, quo.Quo, quo.Rem)
			}

		case <-t.C:
			fmt.Println("10s timeout!")
			return
		}
	}

	//3.打印结果
	fmt.Printf("Divide: %d/%d=%d...%d\n", args.A, args.B, quo.Quo, quo.Rem)
}
