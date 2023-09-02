package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"

	localmetadata "github.com/dushaoshuai/go-usage-examples/grpc/metadata"
	"github.com/dushaoshuai/go-usage-examples/grpc/name_resolution_and_load_balancing/client_side_load_balancing/common"
	"github.com/dushaoshuai/go-usage-examples/grpc/sayhello"
)

var (
	DNSDiscoveryClient sayhello.HelloClient
)

func init() {
	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
	}
	target := fmt.Sprintf("%s:///%s:%d", "dns", "grpc-client-side-lb-headless-svc", common.SayHelloSvcPort)
	conn, err := grpc.Dial(target, dialOpts...)
	if err != nil {
		panic(err)
	}
	DNSDiscoveryClient = sayhello.NewHelloClient(conn)
}

type sayHelloResp struct {
	Msg    string
	Server string // server IP
	Error  string
}

func DNSDiscoverySayHello(name string) sayHelloResp {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var header metadata.MD
	resp, err := DNSDiscoveryClient.SayHello(ctx, &sayhello.HelloReq{Name: name}, grpc.Header(&header))
	return sayHelloResp{
		Msg:    resp.GetMessage(),
		Server: localmetadata.GetServerIP(header),
		Error: func() string {
			if err != nil {
				return err.Error()
			}
			return "nil"
		}(),
	}
}
