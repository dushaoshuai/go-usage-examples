package etcd_test

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"time"

	"go.etcd.io/etcd/client/v3"
)

func ExampleGetWithPrefix() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		if err := cli.Close(); err != nil {
			logrus.Error(err)
		}
	}()

	mustPut(cli, context.TODO(), "/etc/X11/xorg.conf.d/00-keyboard.conf", "ok")
	mustPut(cli, context.TODO(), "/etc/dkms/framework.conf", "ok")
	mustPut(cli, context.TODO(), "/etc/fonts/fonts.conf", "ok")
	mustPut(cli, context.TODO(), "/etc/fonts/conf.d/10-hinting-slight.conf", "ok")
	mustPut(cli, context.TODO(), "/etc/tlp.conf", "ok")

	getResp := mustGet(cli, context.TODO(), "/etc/", clientv3.WithPrefix())
	for _, kv := range getResp.Kvs {
		fmt.Printf("%s: %s\n", kv.Key, kv.Value)
	}
	// Output:
	// /etc/X11/xorg.conf.d/00-keyboard.conf: ok
	// /etc/dkms/framework.conf: ok
	// /etc/fonts/conf.d/10-hinting-slight.conf: ok
	// /etc/fonts/fonts.conf: ok
	// /etc/tlp.conf: ok
}
