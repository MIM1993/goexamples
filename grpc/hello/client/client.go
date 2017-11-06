package main

import (
	pb "github.com/fengchunjian/goexamples/grpc/hello/proto" //引入proto包
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	c := pb.NewHelloClient(conn)
	reqBody := new(pb.HelloRequest)
	reqBody.Name = "gRPC"
	r, err := c.SayHello(context.Background(), reqBody)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(r.Message)
}
