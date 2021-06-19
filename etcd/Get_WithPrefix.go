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
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	mustPut(cli, context.TODO(), "/etc/X11/xorg.conf.d/00-keyboard.conf", "ok")
	mustPut(cli, context.TODO(), "/etc/dkms/framework.conf", "ok")
	mustPut(cli, context.TODO(), "/etc/fonts/fonts.conf", "ok")
	mustPut(cli, context.TODO(), "/etc/fonts/conf.d/10-hinting-slight.conf", "ok")
	mustPut(cli, context.TODO(), "/etc/tlp.conf", "ok")

	getResp, err := cli.Get(context.TODO(), "/etc/", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range getResp.Kvs {
		fmt.Printf("%s : %s\n", kv.Key, kv.Value)
	}
}

func mustPut(c *clientv3.Client, ctx context.Context, key, val string, opts ...clientv3.OpOption) *clientv3.PutResponse {
	putResp, err := c.Put(ctx, key, val, opts...)
	if err != nil {
		log.Fatal(err)
	}
	return putResp
}
