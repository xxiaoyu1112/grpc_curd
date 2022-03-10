package main

import (
	"create_data/idl/create_data"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50001"
)

func main() {

	// 监听本地的 port 端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer() // 创建gRPC服务器

	create_data.RegisterCreateDataCollectServer(s, &Server{})

	log.Printf("server listening at %v", lis.Addr())

	// Serve 方法在 lis 上接受传入连接，为每个连接创建一个ServerTransport和server的goroutine。
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应它们。
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
