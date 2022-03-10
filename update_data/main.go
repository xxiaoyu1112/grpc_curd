package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	"update_data/idl/update_data"
)

const (
	port = ":50003"
)

func main() {

	// 1. 监听本地的 port 端口,并打印监听端口
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Printf("Server listening at %v", lis.Addr())

	// 2. 创建grpc服务器
	s := grpc.NewServer()

	// 3. 注册服务
	update_data.RegisterUpdateDataManageServer(s, &Server{})

	// 4. 创建监听
	s.Serve(lis)
}
