package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Go")
		time.Sleep(3 * time.Second)
	}
}
