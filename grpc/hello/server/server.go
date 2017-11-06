package main

import (
	pb "github.com/fengchunjian/goexamples/grpc/hello/proto" //引入编译生成的包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
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

	s := grpc.NewServer()                   //实例化grpc Server
	pb.RegisterHelloServer(s, HelloService) //注册HelloService

	//grpclog.Println("Listen on " + Address)
	log.Println("Listen on " + Address)
	s.Serve(listen)
}
