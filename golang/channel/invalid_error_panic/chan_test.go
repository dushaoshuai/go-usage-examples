package invalid_error_panic_test

// It is an error if ch is a receive-only channel.
func Example_close_a_receive_only_chan() {
	// invalid operation: cannot close receive-only channel ch (variable of type <-chan int)
	// var ch <-chan int = make(chan int)
	// close(ch)

	// Output:
}

func Example_send_to_a_receive_only_chan() {
	// invalid operation: cannot send to receive-only channel ch (variable of type <-chan int)
	// var ch <-chan int = make(chan int)
	// ch <- 8

	// Output:
}

func Example_receive_from_a_send_only_chan() {
	// invalid operation: cannot receive from send-only channel ch (variable of type chan<- int)
	// var ch chan<- int = make(chan int)
	// <-ch

	// Output:
}
