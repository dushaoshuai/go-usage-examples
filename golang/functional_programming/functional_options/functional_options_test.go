package functional_options_test

import (
	"fmt"

	"github.com/dushaoshuai/go-usage-examples/golang/functional_programming/functional_options/functional_options"
)

func Example_functional_options_pattern1() {
	h1 := functional_options.NewHouse()
	h2 := functional_options.NewHouse(
		functional_options.WithConcrete(),
		functional_options.WithFloors(10),
		functional_options.WithoutFireplace(),
	)
	fmt.Printf("%+v\n", h1)
	fmt.Printf("%+v\n", h2)

	// Output:
	// &{material:wood hasFireplace:true floors:2}
	// &{material:concrete hasFireplace:false floors:10}
}

func Example_functional_options_pattern2() {
	s1 := functional_options.NewServer()
	s2 := functional_options.NewServer(
		functional_options.ReadBufferSize(10),
		functional_options.InitialConnWindowSize(10),
		functional_options.WriteBufferSize(10),
		functional_options.InitialWindowSize(10),
	)
	fmt.Printf("%+v\n", s1)
	fmt.Printf("%+v\n", s2)

	// Output:
	// &{opts:{maxReceiveMessageSize:4194304 maxSendMessageSize:2147483647 writeBufferSize:32768 readBufferSize:32768 connectionTimeout:120000000000 initialWindowSize:0 initialConnWindowSize:0}}
	// &{opts:{maxReceiveMessageSize:4194304 maxSendMessageSize:2147483647 writeBufferSize:10 readBufferSize:10 connectionTimeout:120000000000 initialWindowSize:10 initialConnWindowSize:10}}
}
