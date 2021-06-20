package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	grantResp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	key := "foo"
	_, err = cli.Put(context.TODO(), key, "bar", clientv3.WithLease(grantResp.ID))
	if err != nil {
		log.Fatal(err)
	}

	// the key "foo" will be kept forever
	chanKeepAliveResp, err := cli.KeepAlive(context.TODO(), grantResp.ID)
	if err != nil {
		log.Fatal(err)
	}

	// watch keepAlive status
	done := make(chan struct{})
	go func() {
		for keepAliveResp := range chanKeepAliveResp {
			log.Printf("ttl for key %s: %v", key, keepAliveResp.TTL)

		}
		select {
		case <-done:
		default:
			err := errors.New(fmt.Sprintf("the underlying keep alive stream for key %s is interrupted in some way the client cannot handle itself, e.g. context \"ctx\" is canceled or timed out", key))
			panic(err)
		}
	}()

	// sleep 10 seconds to simulate program running
	time.Sleep(10 * time.Second)
	// now it's time to end this program
	close(done)
}

// $ go run Lease_With_KeepAlive.go
// 2021/06/20 09:35:09 ttl for key foo: 5
// 2021/06/20 09:35:11 ttl for key foo: 5
// 2021/06/20 09:35:13 ttl for key foo: 5
// 2021/06/20 09:35:15 ttl for key foo: 5
// 2021/06/20 09:35:17 ttl for key foo: 5
