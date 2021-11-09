package etcd_test

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

	// watch withPrefix withPrevKV
	chanWatchResp := cli.Watch(context.TODO(), "foo", clientv3.WithPrefix(), clientv3.WithPrevKV())
	go func() {
		for watchResp := range chanWatchResp {
			// Canceled is used to indicate watch failure.
			// If the watch failed and the stream was about to close, before the channel is closed,
			// the channel sends a final response that has Canceled set to true with a non-nil Err().
			if watchResp.Canceled {
				log.Print("watch failed and the stream was about to close", watchResp.Err())
				break
			}
			for _, ev := range watchResp.Events {
				if ev.PrevKv == nil { // put
					fmt.Printf("%s %s : %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				} else if ev.PrevKv != nil && ev.Kv.Value == nil { //delete, clientv3.WithPreKV() only works with delete operation
					fmt.Printf("%s %s, previous %s : %s\n", ev.Type, ev.Kv.Key, ev.PrevKv.Key, ev.PrevKv.Value)
				}
			}
		}
	}()

	// put
	keys := []string{"foo1", "foo2", "foo3", "foo4", "foo5"}
	vals := []string{"bar1", "bar2", "bar3", "bar4", "bar5"}
	for i := range keys {
		if _, err = cli.Put(context.TODO(), keys[i], vals[i]); err != nil {
			log.Fatal(err)
		}
	}
	// delete
	time.Sleep(3 * time.Second)
	if _, err = cli.Delete(context.TODO(), "foo", clientv3.WithPrefix()); err != nil {
		log.Fatal(err)
	}

	// put withLease
	grantResp, err := cli.Grant(context.TODO(), 10)
	if err != nil {
		log.Fatal(err)
	}
	if _, err = cli.Put(context.TODO(), "foo6", "bar6", clientv3.WithLease(grantResp.ID)); err != nil {
		log.Fatal(err)
	}
	time.Sleep(15 * time.Second)
}
