package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"mygo/helloworld/proto"
)

var (
	ip   = flag.String("ip", "127.0.0.1", "")
	port = flag.Int("port", 50001, "")
)

func main() {
	flag.Parse()
	addr := fmt.Sprintf("%v:%v", *ip, *port)

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalln(err)
		return
	}
	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	sayHello(client)
}

func getHelloRequest() *proto.HelloRequest {
	req := &proto.HelloRequest{
		Name:     "david",
		Gender:   proto.Gender_MALE,
		Age:      20,
		Birthday: 1991,
		Addr: &proto.Address{
			Province: "China",
			City:     "BEIJING",
		},
	}
	return req
}

func sayHello(client proto.GreeterClient) {
	ctx := context.Background()
	req := getHelloRequest()
	resq, err := client.SayHello(ctx, req)
	if err != nil {
		log.Fatal(err)
		return
	}

	log.Println(resq.Msg)
}
