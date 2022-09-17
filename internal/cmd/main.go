package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"maxblog-be-template/src/pb"
	"maxblog-be-template/src/service"
	"net"
)

func main() {
	ip := flag.String("ip", "127.0.0.1", "输入ip")
	port := flag.Int("port", 9095, "输入port")
	flag.Parse()
	addr := fmt.Sprintf("%s:%d", *ip, *port)

	const ConfigDir = "env/raw/dev.yaml"
	ctx := context.Background()

	server := grpc.NewServer()
	pb.RegisterDataServiceServer(server, &service.DataServer{})
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	err = server.Serve(listen)
	if err != nil {
		panic(err)
	}
}
