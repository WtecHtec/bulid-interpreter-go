package main

import (
	"fmt"
	pb "monkey/test/protoc_demo/proto/hello"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50052"
)

// 定义helloService并实现约定的接口
type helloService struct {
	pb.UnimplementedHelloServer
}

// HelloService Hello服务
var HelloService = helloService{}

// SayHello 实现Hello服务接口
func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

func main() {
	// grpc服务
	// listen, err := net.Listen("tcp", Address)
	// if err != nil {
	// 	grpclog.Fatalf("Failed to listen: %v", err)
	// }

	// // 实例化grpc Server
	// s := grpc.NewServer()

	// // 注册HelloService
	// pb.RegisterHelloServer(s, HelloService)

	// fmt.Println("Listen on " + Address)
	// s.Serve(listen)

	// GATEWAY 使用
	// 1. 启动GRPC服务
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	pb.RegisterHelloServer(s, HelloService)
	// Serve gRPC Server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50052))
	if err != nil {
		fmt.Println("GRPC failed to listen: %v", err)
	}
	fmt.Println("Serving gRPC on 0.0.0.0" + fmt.Sprintf(":%d", 50052))
	go func() {
		if err := s.Serve(lis); err != nil {
			fmt.Println("failed to serve: %v", err)
		}
	}()

	// 2. 启动 HTTP 服务
	// Create a client connection to the gRPC server we just started
	// This is where the gRPC-Gateway proxies the requests
	conn, err := grpc.Dial(
		Address,
		grpc.WithInsecure(),
	)
	if err != nil {
		fmt.Println("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// Register Greeter
	err = pb.RegisterHelloHandler(context.Background(), gwmux, conn)
	if err != nil {
		fmt.Println("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8099),
		Handler: gwmux,
	}

	fmt.Println("Serving gRPC-Gateway on http://0.0.0.0" + fmt.Sprintf(":%d", 8099))
	fmt.Println(gwServer.ListenAndServe())
}
