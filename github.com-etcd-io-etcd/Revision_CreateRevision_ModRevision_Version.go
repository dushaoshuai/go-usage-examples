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

	keys := []string{"foo1", "foo2", "foo3", "foo4", "foo5"}
	vals := []string{"bar1", "bar2", "bar3", "bar4", "bar5"}
	// put
	for i := range keys {
		putResp, err := cli.Put(context.TODO(), keys[i], vals[i])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("put key %s, reversion number: %v\n", keys[i], putResp.Header.Revision)
	}
	fmt.Println()
	// get
	getResp, err := cli.Get(context.TODO(), "foo", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range getResp.Kvs {
		fmt.Printf("get key %s, createRevison number: %v, modRevison number: %v, version: %v\n", kv.Key, kv.CreateRevision, kv.ModRevision, kv.Version)
	}
	fmt.Println()
	// put
	for i := range keys {
		putResp, err := cli.Put(context.TODO(), keys[i], vals[i])
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("put key %s, reversion number: %v\n", keys[i], putResp.Header.Revision)
	}
	fmt.Println()
	// get
	getResp, err = cli.Get(context.TODO(), "foo", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range getResp.Kvs {
		fmt.Printf("get key %s, createRevison number: %v, modRevison number: %v, version: %v\n", kv.Key, kv.CreateRevision, kv.ModRevision, kv.Version)
	}
	fmt.Println()
	// delete
	delResp, err := cli.Delete(context.TODO(), "foo", clientv3.WithPrefix(), clientv3.WithPrevKV())
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range delResp.PrevKvs {
		fmt.Printf("delete key %s, createRevison number: %v, modRevison number: %v, version: %v\n", kv.Key, kv.CreateRevision, kv.ModRevision, kv.Version)
	}
}

// revision is the key-value store revision when the request was applied.
// For watch progress responses, the header.revision indicates progress. All future events
// recieved in this stream are guaranteed to have a higher revision number than the
// header.revision number.

// clientv3.ResponseHead.Reversion auto increments by 1
// every time there's an put/delete operation. I'm not sure
// if that applies to all operations in a clientv3.Txn.

// The CreateReversion of a key is the Reversion when the key in inserted.
// The ModReversion of a key is the Reversion when the key was inserted or updated.
// The Version of a key auto increments by 1 every time it's updated, starts from 1 when the key was inserted.
// When a key was deleted, it does not exists any more, it's a new key with Version 1 next time it's inserted.

// $ go run etcd/CreateRevison_ModReviosn.go
// put key foo1, reversion number: 205
// put key foo2, reversion number: 206
// put key foo3, reversion number: 207
// put key foo4, reversion number: 208
// put key foo5, reversion number: 209
//
// get key foo1, createRevison number: 205, modRevison number: 205, version: 1
// get key foo2, createRevison number: 206, modRevison number: 206, version: 1
// get key foo3, createRevison number: 207, modRevison number: 207, version: 1
// get key foo4, createRevison number: 208, modRevison number: 208, version: 1
// get key foo5, createRevison number: 209, modRevison number: 209, version: 1
//
// put key foo1, reversion number: 210
// put key foo2, reversion number: 211
// put key foo3, reversion number: 212
// put key foo4, reversion number: 213
// put key foo5, reversion number: 214
//
// get key foo1, createRevison number: 205, modRevison number: 210, version: 2
// get key foo2, createRevison number: 206, modRevison number: 211, version: 2
// get key foo3, createRevison number: 207, modRevison number: 212, version: 2
// get key foo4, createRevison number: 208, modRevison number: 213, version: 2
// get key foo5, createRevison number: 209, modRevison number: 214, version: 2
//
// delete key foo1, createRevison number: 205, modRevison number: 210, version: 2
// delete key foo2, createRevison number: 206, modRevison number: 211, version: 2
// delete key foo3, createRevison number: 207, modRevison number: 212, version: 2
// delete key foo4, createRevison number: 208, modRevison number: 213, version: 2
// delete key foo5, createRevison number: 209, modRevison number: 214, version: 2
