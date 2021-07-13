// https://github.com/etcd-io/etcd/tree/main/client/v3#get-started

// More code examples can be found at :
// https://github.com/etcd-io/etcd/tree/main/tests/integration/clientv3/examples
// More code examples about error handling can be found at :
// https://pkg.go.dev/go.etcd.io/etcd/client/v3#section-documentation

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	// Clients are safe for concurrent use by multiple goroutines.
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// etcd v3 uses gRPC for remote procedure calls.
	// And clientv3 uses grpc-go to connect to etcd.
	// Make sure to close the client after using it.
	// If the client is not closed, the connection will have leaky goroutines.
	// To specify client request timeout, pass context.WithTimeout to APIs:

	// put
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	_, err = cli.Put(ctx, "love", "xiaoying")
	if err != nil {
		log.Fatal(err)
	}
	_, err = cli.Put(context.TODO(), "lo", "xiao")
	if err != nil {
		log.Fatal(err)
	}
	// get
	getResp, err := cli.Get(ctx, "lo", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}
	// https://github.com/etcd-io/etcd/blob/main/tests/integration/clientv3/examples/example_kv_test.go
	for _, kv := range getResp.Kvs {
		// https://pkg.go.dev/go.etcd.io/etcd/api/v3@v3.5.0/mvccpb#KeyValue
		fmt.Printf("%s : %s\n", kv.Key, kv.Value)
	}

	// The grpc load balancer is registered statically and is shared across etcd clients.
	// To enable detailed load balancer logging,
	// set the ETCD_CLIENT_DEBUG environment variable. E.g. "ETCD_CLIENT_DEBUG=1".
}

//$ go run getStart.go
// lo : xiao
// love : xiaoying
// $
// $ go run getStart.go
// lo : xiao
// love : xiaoying
// $
// $ go run getStart.go
// lo : xiao
// love : xiaoying