package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"mygo/helloworld/proto"
	"net"
)

var (
	port = flag.Int("port", 50001, "")
)

type server struct {
	proto.UnimplementedGreeterServer //UnimplementedGreeterServer 实现了 GreeterServer 接口
}

func (server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReplay, error) {
	log.Printf("server recv: %v", req)
	return &proto.HelloReplay{
		Msg: "hello client",
	}, nil
}
func (server) SayHelloClientStream(stream proto.Greeter_SayHelloClientStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloClientStream not implemented")
}
func (server) SayHelloServerStream(req *proto.HelloRequest, stream proto.Greeter_SayHelloServerStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloServerStream not implemented")
}
func (server) SayHelloTwoWayStream(req *proto.HelloRequest, stream proto.Greeter_SayHelloTwoWayStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method SayHelloTwoWayStream not implemented")
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	log.Println(fmt.Sprintf("server listen  :%d", *port))
	err = s.Serve(lis)

	if err != nil {
		log.Fatal(err)
	}
}
