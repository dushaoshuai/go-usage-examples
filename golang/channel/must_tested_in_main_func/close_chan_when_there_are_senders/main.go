package main

import (
	"sync"

	"github.com/dushaoshuai/goloop"
)

func main() {
	ch := make(chan int)

	var wg sync.WaitGroup
	for range goloop.Repeat(3) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- 0
		}()
	}

	close(ch) // Note: if there are still senders, they will panic after the channel is closed
	wg.Wait()

	// Output:
	// panic: send on closed channel
	//
	// goroutine 8 [running]:
	// main.main.func1()
	//        /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/channel/must_tested_in_main_func/close_chan_when_there_are_senders/main.go:17 +0x65
	// created by main.main
	//        /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/channel/must_tested_in_main_func/close_chan_when_there_are_senders/main.go:15 +0x52
	// panic: send on closed channel
	//
	// goroutine 6 [running]:
	// main.main.func1()
	//        /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/channel/must_tested_in_main_func/close_chan_when_there_are_senders/main.go:17 +0x65
	// created by main.main
	//        /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/channel/must_tested_in_main_func/close_chan_when_there_are_senders/main.go:15 +0x52
}
