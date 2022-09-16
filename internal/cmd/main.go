package main

import (
	"flag"
	"fmt"
	"github.com/golang/protobuf/protoc-gen-go/grpc"
	"net"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "输入ip")
	port := flag.Int("port", 9095, "输入port")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *ip, *port)
	server := grpc.NewServer()
	pb.RegisterAccountServiceServer(server, &biz.AccountServer{})
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
