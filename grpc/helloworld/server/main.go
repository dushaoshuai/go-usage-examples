package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/dushaoshuai/go-usage-examples/grpc/helloworld/proto"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "the server port")

type server struct {
	proto.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *proto.HelloReq) (*proto.HelloResp, error) {
	log.Printf("Received: %v", in.GetName())
	return &proto.HelloResp{Message: "Hello " + in.GetName()}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
