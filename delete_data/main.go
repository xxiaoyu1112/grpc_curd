package main

import (
	"delete_data/idl/delete_data"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50004"
)

func main() {

	// 1. 监听本地的 port 端口,并打印监听端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("server listening at %v", lis.Addr())

	// 2. 创建grpc服务器
	s := grpc.NewServer()

	// 3. 在gRPC服务端注册服务
	delete_data.RegisterDeleteDataManageServer(s, &Server{})

	// 4.创建监听
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
