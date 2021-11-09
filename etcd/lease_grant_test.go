package etcd_test

import (
	"context"
	"log"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func ExampleLeaseGrant() {
	log.SetFlags(0)

	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := cli.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// minimum lease TTL is 5-second ???
	grantResp, err := cli.Grant(context.TODO(), 1)
	if err != nil {
		log.Fatal(err)
	}

	// after 5 seconds, the key "foo" will be removed
	key1 := "foo"
	_, err = cli.Put(context.TODO(), key1, "bar", clientv3.WithLease(grantResp.ID))
	if err != nil {
		log.Fatal(err)
	}

	// after 5 seconds, the key "love" will be removed
	key2 := "love"
	_, err = cli.Put(context.TODO(), key2, "true", clientv3.WithLease(grantResp.ID))
	if err != nil {
		log.Fatal(err)
	}
	// Output:
}
