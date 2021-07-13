package main

import (
	"context"
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

	grantResp, err := cli.Grant(context.TODO(), 10)
	if err != nil {
		log.Fatal(err)
	}

	key := "foo"
	_, err = cli.Put(context.TODO(), key, "bar", clientv3.WithLease(grantResp.ID))
	if err != nil {
		log.Fatal(err)
	}

	_, err = cli.Revoke(context.TODO(), grantResp.ID)
	if err != nil {
		log.Fatal(err)
	}

	getResp, err := cli.Get(context.TODO(), key)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("number of keys:", len(getResp.Kvs))
}

// $ go run Lease_Revoke.go
// number of keys: 0
