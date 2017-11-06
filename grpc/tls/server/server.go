package main

import (
	pb "github.com/fengchunjian/goexamples/grpc/tls/proto" //引入编译生成的包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials" //引入grpc认证包
	"log"
	"net"
)

const (
	// gRPC服务地址
	Address = "127.0.0.1:50052"
)

//定义helloService并实现约定的接口
type helloService struct{}

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "Hello " + in.Name + "."
	return resp, nil
}

var HelloService = helloService{}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("failed to listen:%v", err)
	}

	//TLS认证
	creds, err := credentials.NewServerTLSFromFile("../keys/server.pem", "../keys/server.key")
	if err != nil {
		log.Fatalf("failed to generate credentials %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))  //实例化grpc Server，并开启TLS认证
	pb.RegisterHelloServer(s, HelloService) //注册HelloService

	log.Println("Listen on " + Address + " with TLS")
	s.Serve(listen)
}
