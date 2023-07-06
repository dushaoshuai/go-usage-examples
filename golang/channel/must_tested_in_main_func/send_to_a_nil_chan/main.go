package main

func main() {
	var ch chan int // nil channel

	go func() {
		ch <- 0 // A send on a nil channel blocks forever.
	}()

	<-ch // Receiving from a nil channel blocks forever.

	// Output:
	// $ go run main.go
	// fatal error: all goroutines are asleep - deadlock!
	//
	// goroutine 1 [chan receive (nil chan)]:
	// main.main()
	//        /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/channel/must_tested_in_main_func/send_to_a_nil_chan/main.go:10 +0x45
	//
	// goroutine 5 [chan send (nil chan)]:
	// main.main.func1()
	//        /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/channel/must_tested_in_main_func/send_to_a_nil_chan/main.go:7 +0x25
	// created by main.main
	//        /home/shaouai/dev/github.com/dushaoshuai/go-usage-examples/golang/channel/must_tested_in_main_func/send_to_a_nil_chan/main.go:6 +0x3c
	// exit status 2
}
