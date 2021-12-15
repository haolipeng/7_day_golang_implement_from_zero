package main

import (
	"fmt"
	"reflect"
)

type HelloService interface {
	SayHello(request string, reply *string) error
}

type hello struct {
	endpoint  string
	FuncField func()
}

//golang常用的一种写法，确保结构体实现某个接口
var _ HelloService = hello{}

func (h hello) SayHello(request string, reply *string) error {
	//简单回复下
	*reply = "hello:" + request
	return nil
}

func PrintFieldName(val interface{}) {
	t := reflect.TypeOf(val)  //获取对象的类型信息，如该类型(struct)有啥字段，字段是什么类型
	v := reflect.ValueOf(val) //获取对象的运行时表示，例如有啥字段，字段的值是啥

	fields := t.NumField()
	for i := 0; i < fields; i++ {
		field := t.Field(i)
		fieldVal := v.Field(i)

		if fieldVal.CanSet() {
			fmt.Printf("%s 可以被设置", field.Name)
		}

		fmt.Printf("field:%s,field value:%s \n", field.Name, fieldVal.String())
	}
}

//SetFuncField 尝试篡改字段内容
func SetFuncField(val interface{}) {
	tkind := reflect.TypeOf(val).Kind()
	if tkind == reflect.Struct {
		fmt.Println("对象的类型信息")
	}

	//判断值运行时的kind类型
	kind := reflect.ValueOf(val).Kind()
	if kind == reflect.Ptr {
		fmt.Println("指针类型")
	}

	//指针的反射
	v := reflect.ValueOf(val)

	//拿到了指针指向的结构体
	e := v.Elem()

	//拿到了指针指向的结构体的类型信息
	t := e.Type()

	//获取类型的字段信息
	numFields := t.NumField()
	for i := 0; i < numFields; i++ {
		field := t.Field(i)

		//用指针指向的结构体来访问
		fieldval := e.Field(i)

		fmt.Println(field.Name)

		//判断field是否可设置
		if fieldval.CanSet() {
			fmt.Printf("%s 可以被设置\n", field.Name)
			//fieldval.Set(reflect.ValueOf(func() {
			//	fmt.Println("你在调用方法" + field.Name)
			//}))
		}
	}
}

func main() {
	h := &hello{
		endpoint: string("http://localhost:8080/"),
		FuncField: func() {
			fmt.Println("I am FuncField,OK!")
		},
	}

	//PrintFieldName(h)
	h.FuncField() //befor
	SetFuncField(h)
	h.FuncField() //after

	fmt.Println(h)
}
