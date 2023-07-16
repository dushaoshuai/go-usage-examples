package net_test

import (
	"fmt"
	"io"
	"net"
	"os"
	"time"

	"github.com/dushaoshuai/goloop"
)

func Example_tcp_server() {
	go func() {
		time.Sleep(2 * time.Second)
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil {
			panic(err)
		}
		for i := range goloop.Repeat(10) {
			fmt.Fprintf(conn, "Hello %d!\n", i)
		}
		conn.Close()
	}()

	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			panic(err)
		}
		conn.Write()
		io.Copy(os.Stdout, conn)
	}
	// Output:
}
