// Package main implements a server for Greeter service.
package main

import (
	"context"
	"flag"
	"fmt"
	pb "gRpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

//定义port
var (
	port = flag.Int("port", 50052, "The server port")
)

// server用于实现helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello 实现 proto.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	// 监听本地端口
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 新建gRPC服务器实例
	s := grpc.NewServer()
	// 在gRPC服务器注册我们的服务
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	//用服务器 Serve() 方法以及我们的端口信息区实现阻塞等待，直到进程被杀死或者 Stop() 被调用
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
