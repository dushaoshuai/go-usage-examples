package net_test

import (
	"fmt"
	"net"
)

func ExampleResolveIPAddr() {
	ipAddr, err := net.ResolveIPAddr("ip4:1", "135.181.27.174")
	if err != nil {
		panic(err)
	}
	fmt.Println(ipAddr)

	// Output:
	// 135.181.27.174
}
