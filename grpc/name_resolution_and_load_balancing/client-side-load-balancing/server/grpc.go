package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/dushaoshuai/go-usage-examples/grpc/name_resolution_and_load_balancing/client-side-load-balancing/common"
	"github.com/dushaoshuai/go-usage-examples/grpc/sayhello"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", common.SayHelloSvcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	sayhello.RegisterHelloServer(s, helloServer{})
	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
