package server

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quetient struct {
	Quo, Rem int
}

type Arith int

func (a *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quetient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B

	return nil
}

func main() {
	//注册rpc对象
	arith := new(Arith)
	rpc.Register(arith)
	//rpc.RegisterName()

	rpc.HandleHTTP()
	l, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", e)
	}

	go http.Serve(l, nil)
}
