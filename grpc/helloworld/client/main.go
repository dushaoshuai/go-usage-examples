package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/dushaoshuai/go-usage-examples/grpc/helloworld/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const defaultName = "world"

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
	name = flag.String("name", defaultName, "name to greet")
)

func main() {
	flag.Parse()

	// set up a connection to the server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	// contact the server and print out its response
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &proto.HelloReq{Name: *name})
	if err != nil {
		panic(err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
