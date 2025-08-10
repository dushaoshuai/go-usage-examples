package goleak_test

import (
	"sync"
	"testing"

	"go.uber.org/goleak"
)

func Test_goroutine_leak_blocked_chan_send(t *testing.T) {
	// $ go test -run Test_goroutine_leak_blocked_chan_send
	// --- FAIL: Test_goroutine_leak_blocked_chan_send (0.44s)
	//    goleak_test.go:36: found unexpected goroutines:
	//        [Goroutine 7 in state chan send, with github.com/dushaoshuai/go-usage-examples/golang/goleak_test.Test_goroutine_leak_blocked_chan_send.func1 on top of the stack:
	//        goroutine 7 [chan send]:
	//        github.com/dushaoshuai/go-usage-examples/golang/goleak_test.Test_goroutine_leak_blocked_chan_send.func1()
	//                /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/goleak/goleak_test.go:32 +0x3f
	//        created by github.com/dushaoshuai/go-usage-examples/golang/goleak_test.Test_goroutine_leak_blocked_chan_send
	//                /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/goleak/goleak_test.go:30 +0xed
	//        ]
	// FAIL
	// exit status 1
	// FAIL    github.com/dushaoshuai/go-usage-examples/golang/goleak  0.447s
	defer goleak.VerifyNone(t)

	c := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		wg.Done()
		c <- 0
	}()

	wg.Wait()
}

func Test_goroutine_leak_blocked_chan_receive(t *testing.T) {
	// $ go test -run Test_goroutine_leak_blocked_chan_receive
	// --- FAIL: Test_goroutine_leak_blocked_chan_receive (0.45s)
	//    goleak_test.go:64: found unexpected goroutines:
	//        [Goroutine 7 in state chan receive, with github.com/dushaoshuai/go-usage-examples/golang/goleak_test.Test_goroutine_leak_blocked_chan_receive.func1 on top of the stack:
	//        goroutine 7 [chan receive]:
	//        github.com/dushaoshuai/go-usage-examples/golang/goleak_test.Test_goroutine_leak_blocked_chan_receive.func1()
	//                /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/goleak/goleak_test.go:60 +0x3a
	//        created by github.com/dushaoshuai/go-usage-examples/golang/goleak_test.Test_goroutine_leak_blocked_chan_receive
	//                /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/goleak/goleak_test.go:58 +0xed
	//        ]
	// FAIL
	// exit status 1
	// FAIL    github.com/dushaoshuai/go-usage-examples/golang/goleak  0.449s
	defer goleak.VerifyNone(t)

	c := make(chan int)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		wg.Done()
		<-c
	}()

	wg.Wait()
}

func Test_goroutine_leak_loop_forever(t *testing.T) {
	// $ go test -run Test_goroutine_leak_loop_forever
	// --- FAIL: Test_goroutine_leak_loop_forever (0.44s)
	//    goleak_test.go:78: found unexpected goroutines:
	//        [Goroutine 7 in state runnable, with github.com/dushaoshuai/go-usage-examples/golang/goleak_test.Test_goroutine_leak_loop_forever.func1 on top of the stack:
	//        goroutine 7 [runnable]:
	//        github.com/dushaoshuai/go-usage-examples/golang/goleak_test.Test_goroutine_leak_loop_forever.func1()
	//                /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/goleak/goleak_test.go:73 +0x26
	//        created by github.com/dushaoshuai/go-usage-examples/golang/goleak_test.Test_goroutine_leak_loop_forever
	//                /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/goleak/goleak_test.go:71 +0xbc
	//        ]
	// FAIL
	// exit status 1
	// FAIL    github.com/dushaoshuai/go-usage-examples/golang/goleak  0.446s
	defer goleak.VerifyNone(t)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		for {
		}
	}()

	wg.Wait()
}

func Test_no_leak(t *testing.T) {
	// $ go test -run Test_no_leak
	// PASS
	// ok      github.com/dushaoshuai/go-usage-examples/golang/goleak  0.002s
	defer goleak.VerifyNone(t)
}

func Example_run_out_of_memory() {
	for {
		go func() {
			make(chan string) <- "test run out of memory"
		}()
	}

	// Output:
}
