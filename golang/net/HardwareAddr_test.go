package net_test

import (
	"fmt"
	"net"
)

func ExampleHardwareAddr() {
	macAddr, err := net.ParseMAC("98:01:a7:b3:c6:5f")
	if err != nil {
		panic(err)
	}
	fmt.Println(macAddr)
	// Output:
	// 98:01:a7:b3:c6:5f
}
