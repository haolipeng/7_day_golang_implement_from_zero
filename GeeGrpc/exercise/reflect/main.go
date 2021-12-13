package main

import (
	"fmt"
	"reflect"
)

type HelloService interface {
	SayHello(request string, reply *string) error
}

type hello struct {
	endpoint string
}

//golang常用的一种写法，确保结构体实现某个接口
var _ HelloService = hello{}

func (h hello) SayHello(request string, reply *string) error {
	//简单回复下
	*reply = "hello:" + request
	return nil
}

func (h hello) SayGoodBye(request string, reply *string) error {
	*reply = "goodbye:" + request
	return nil
}

func PrintFuncName(v interface{}) {
	//TypeOf returns the reflection Type that represents the dynamic type of i.
	//获取结构体的所有方法
	t := reflect.TypeOf(v)
	num := t.NumMethod()

	for i := 0; i < num; i++ {
		m := t.Method(i)
		fmt.Println(m.Name)
	}
}

func main() {
	var h = hello{}
	PrintFuncName(h)
}
