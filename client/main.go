package main

import (
	"context"
	"flag"
	pb "gRpc/proto"
	"log"
	"time"

	"google.golang.org/grpc"
	// pb "proto/proto"
	//package path = module path +"/"+ package name
)

const (
	defaultName = "world"
)

var (
	addr = flag.String("addr", "localhost:50052", "the address to connect to")
	name = flag.String("name", defaultName, "Name to greet")
)

func main() {
	flag.Parse()
	// 连接服务器
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
	// 建立gRPC连接

	c := pb.NewGreeterClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// 调用我们的服务(SayHello方法)
	// 同时传入了一个 context.Context ，在有需要时可以让我们改变RPC的行为，比如超时/取消一个正在运行的RPC
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	// 打印返回值
	log.Printf("Greeting: %s", r.GetMessage())
}
