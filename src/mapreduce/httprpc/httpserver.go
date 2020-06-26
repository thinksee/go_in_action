package main

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

type Quotient struct {
	Quo, Rem int
}

type Arith int

// the right way for rpc: func (t *T) MethodName(argType T1, replyType *T2) error
func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("Divide by zero!!!")
	}
	quo.Quo = args.A / args.B
	quo.Quo = args.A * args.B
	return nil
}

func main() {
	// new(T) return (the pointer of T and the value of T will init.)
	// make only be used slice/map/channel, and return the reference of T
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(listen, nil)
}
