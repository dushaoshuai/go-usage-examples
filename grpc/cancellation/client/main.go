package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/dushaoshuai/go-usage-examples/grpc/cancellation/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()

	// set up a connection to the server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewCancelClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	time.AfterFunc(3*time.Second, func() {
		fmt.Println("3 seconds elapsed")
		fmt.Println("cancelling context")
		cancel()
	})

	_, err = c.IWillCancel(ctx, &emptypb.Empty{})
	if code := status.Code(err); code != codes.Canceled {
		log.Printf("got error code %v, want %v", code, codes.Canceled)
	}
}
