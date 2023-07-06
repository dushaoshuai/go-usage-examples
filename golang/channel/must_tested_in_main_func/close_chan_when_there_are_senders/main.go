package main

import (
	"sync"
	"time"

	"github.com/dushaoshuai/goloop"
)

func main() {
	ch := make(chan int)

	var wg1 sync.WaitGroup
	var wg2 sync.WaitGroup
	for range goloop.Repeat(3) {
		wg1.Add(1)
		wg2.Add(1)
		go func() {
			defer wg1.Done()
			wg2.Done()
			ch <- 0
		}()
	}

	wg2.Wait()
	time.Sleep(3 * time.Second) // wait until all three senders are blocked on this channel
	close(ch)                   // Note: if there are still senders, they will panic after the channel is closed
	wg1.Wait()

	// Output:
	// $ go run main.go
	// panic: send on closed channel
	//
	// goroutine 7 [running]:
	// main.main.func1()
	//        /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/channel/must_tested_in_main_func/close_chan_when_there_are_senders/main.go:21 +0x7e
	// created by main.main
	//        /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/channel/must_tested_in_main_func/close_chan_when_there_are_senders/main.go:18 +0x65
	// exit status 2
}
