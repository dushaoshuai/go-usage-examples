package main

import (
	"context"
	"fmt"
	"log"
	"sync"
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

	// minimum lease TTL is 5-second
	grantResp, err := cli.Grant(context.TODO(), 5)
	if err != nil {
		log.Fatal(err)
	}

	// wait for all goroutines to accomplish
	wg := &sync.WaitGroup{}

	// after 5 seconds, the key "foo" will be removed
	key1 := "foo"
	_, err = cli.Put(context.TODO(), key1, "bar", clientv3.WithLease(grantResp.ID))
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	go watchLease(cli, key1, wg)

	// after 5 seconds, the key "dummy" will be removed
	key2 := "love"
	_, err = cli.Put(context.TODO(), key2, "true", clientv3.WithLease(grantResp.ID))
	if err != nil {
		log.Fatal(err)
	}
	wg.Add(1)
	go watchLease(cli, key2, wg)

	wg.Wait()
	fmt.Println("Done!")
}

func watchLease(c *clientv3.Client, key string, wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(1 * time.Second)
	for {
		getResp, err := c.Get(context.TODO(), key)
		if err != nil {
			log.Fatal(err)
		}
		if len(getResp.Kvs) == 0 {
			log.Printf("the key %s has been removed due to lease expiration", key)
			break
		}
		fmt.Printf("%s : %s, %v\n", key, getResp.Kvs[0].Value, time.Now())
		<-ticker.C
	}
	ticker.Stop()
}
