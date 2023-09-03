package metadata

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/dushaoshuai/go-usage-examples/golang/net"
)

var (
	serverIPKey = "server-ip-key"
)

func GetServerIP(md metadata.MD) string {
	IPs := md.Get(serverIPKey)
	if len(IPs) == 0 {
		return ""
	}
	return IPs[0]
}

func SendServerIP(ctx context.Context) {
	header := metadata.Pairs(serverIPKey, net.GetOutboundIP())
	grpc.SendHeader(ctx, header)
}
