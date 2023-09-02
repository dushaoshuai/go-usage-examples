package main

import (
	"context"

	localmetadata "github.com/dushaoshuai/go-usage-examples/grpc/metadata"
	"github.com/dushaoshuai/go-usage-examples/grpc/sayhello"
)

type helloServer struct {
	sayhello.UnimplementedHelloServer
}

func (helloServer) SayHello(ctx context.Context, in *sayhello.HelloReq) (*sayhello.HelloResp, error) {
	localmetadata.SendServerIP(ctx)
	return &sayhello.HelloResp{}, nil
}
