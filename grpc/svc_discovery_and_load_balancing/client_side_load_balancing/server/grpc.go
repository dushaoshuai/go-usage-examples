package main

import (
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"

	"github.com/dushaoshuai/go-usage-examples/grpc/sayhello"
	"github.com/dushaoshuai/go-usage-examples/grpc/svc_discovery_and_load_balancing/client_side_load_balancing/common"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", common.SayHelloSvcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			// gRPC DNS 解析机制：
			// 两次解析之间最少间隔 30 秒，
			// 有连接关闭或失败时才触发解析。
			//
			// 使用 DNS resolver 进行[服务发现]的问题：
			// 1. 两次解析之间的时间间隔较大；
			// 2. 服务端扩容时，无法触发客户端的解析，客户端无法感知服务端新副本的存在；
			// 3. k8s 要使用 Headless Service，造成一定的使用限制；
			//
			// 针对第二个问题的解决方案：
			// 设置连接的最大存活时间为 1 分钟，
			// 连接关闭时，触发客户端进行 DNS 解析。
			// 参考 https://github.com/grpc/grpc-go/issues/3170#issuecomment-552517779
			// 其实这里还是有问题，因为比较难确定一个比较合适的连接最大存活时间。
			MaxConnectionAge: 60 * time.Second,
		}),
	}
	s := grpc.NewServer(opts...)
	sayhello.RegisterHelloServer(s, helloServer{})

	log.Printf("server listening at %v", lis.Addr())
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
