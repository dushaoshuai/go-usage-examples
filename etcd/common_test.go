package etcd_test

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/client/v3"
)

func mustPut(c *clientv3.Client, ctx context.Context, key, val string, opts ...clientv3.OpOption) *clientv3.PutResponse {
	putResp, err := c.Put(ctx, key, val, opts...)
	if err != nil {
		logrus.Fatal(err)
	}
	return putResp
}

func mustGet(c *clientv3.Client, ctx context.Context, key string, opts ...clientv3.OpOption) *clientv3.GetResponse {
	getResp, err := c.Get(ctx, key, opts...)
	if err != nil {
		logrus.Fatal(err)
	}
	return getResp
}
