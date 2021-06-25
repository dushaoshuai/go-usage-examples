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

	// put
	mustPut(cli, context.TODO(), "/etc/X11/xorg.conf.d/00-keyboard.conf", "")
	mustPut(cli, context.TODO(), "/etc/dkms/framework.conf", "")
	mustPut(cli, context.TODO(), "/etc/fonts/fonts.conf", "")
	mustPut(cli, context.TODO(), "/etc/fonts/conf.d/10-hinting-slight.conf", "")
	mustPut(cli, context.TODO(), "/etc/tlp.conf", "")
	// get
	getResp, err := cli.Get(context.TODO(), "/etc/", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range getResp.Kvs {
		fmt.Printf("%s : %s\n", kv.Key, kv.Value)
	}
	fmt.Println()
	// delete
	delResp, err := cli.Delete(context.TODO(), "/etc/", clientv3.WithPrefix(), clientv3.WithPrevKV())
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range delResp.PrevKvs {
		fmt.Printf("deleted %s : %s\n", kv.Key, kv.Value)
	}
	fmt.Println("Done!")
	// deleting k-v that doesn't exists is not an error
	_, err = cli.Delete(context.TODO(), "keyDoesn'tExists")
	if err != nil {
		log.Print(err)
	} else {
		log.Print("delete keyDoesn'tExists successfully!")
	}
}

func mustPut(c *clientv3.Client, ctx context.Context, key, val string, opts ...clientv3.OpOption) *clientv3.PutResponse {
	putResp, err := c.Put(ctx, key, val, opts...)
	if err != nil {
		log.Fatal(err)
	}
	return putResp
}

// $ go run Delete_WithPrefix.go
// /etc/X11/xorg.conf.d/00-keyboard.conf :
// /etc/dkms/framework.conf :
// /etc/fonts/conf.d/10-hinting-slight.conf :
// /etc/fonts/fonts.conf :
// /etc/tlp.conf :
//
// deleted /etc/X11/xorg.conf.d/00-keyboard.conf :
// deleted /etc/dkms/framework.conf :
// deleted /etc/fonts/conf.d/10-hinting-slight.conf :
// deleted /etc/fonts/fonts.conf :
// deleted /etc/tlp.conf :
// Done!
