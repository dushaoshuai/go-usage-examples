package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/dushaoshuai/go-usage-examples/grpc/cancellation/proto"

	"google.golang.org/grpc"
)

var port = flag.Int("port", 50051, "the server port")

type server struct {
	proto.UnimplementedCancelServer
}

func (s *server) IWillCancel(ctx context.Context, _ *emptypb.Empty) (*emptypb.Empty, error) {
	t := time.Tick(1 * time.Second)
	for {
		select {
		case <-ctx.Done():
			log.Println(ctx.Err())
			return nil, ctx.Err()
		case <-t:
			log.Println("processing")
		}
	}
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	proto.RegisterCancelServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		panic(err)
	}
}
