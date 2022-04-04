package net_test

import (
	"fmt"
	"net"
)

func ExampleJoinHostPort_ip4() {
	host := "192.168.188.156"
	port := "7890"
	ip4Addr := net.JoinHostPort(host, port)
	fmt.Println(ip4Addr)
	// Output:
	// 192.168.188.156:7890
}

func ExampleJoinHostPort_ip6() {
	host := "fe80::1f12:9b4d:bc6b:6c2d"
	port := "5670"
	ip6Addr := net.JoinHostPort(host, port)
	fmt.Println(ip6Addr)
	// Output:
	// [fe80::1f12:9b4d:bc6b:6c2d]:5670
}