package etcd_test

import (
	"context"
	"fmt"
	"log"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/concurrency"
)

func main() {
	// client
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	// session
	s, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	// mutex
	m := concurrency.NewMutex(s, "/cron/scanFiles/reload.conf")
	// acquire lock for s
	if err := m.TryLock(context.TODO()); err != nil {
		log.Print(err)
		return
	}
	fmt.Println("acquired lock for session")
	// do something
	timer := time.After(10 * time.Second)
	ticker := time.Tick(1 * time.Second)
outerLoop:
	for {
		fmt.Println("Do something!")
		<-ticker
		select {
		case <-timer:
			break outerLoop
		default:
		}
	}
}
