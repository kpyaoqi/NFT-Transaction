package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"yqnft/ContractManage/service"
	pb "yqnft/proto"
)

var u = service.ContractService{}

func main() {
	addr := "127.0.0.1:7080"
	// 1.监听
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Printf("监听异常:%s\n", err)
	}
	fmt.Printf("监听端口：%s\n", addr)
	// 2.实例化gRPC
	s := grpc.NewServer()
	// 3.在gRPC上注册微服务
	pb.RegisterContractServiceServer(s, &u)
	// 4.启动服务端
	s.Serve(listener)
}
