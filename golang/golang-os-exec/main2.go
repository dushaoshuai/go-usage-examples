package main

import (
	"context"
	"log"
	"os/exec"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
		log.Fatal(err)
	}
}

// $ go run main2.go
// 2021/06/17 19:30:45 signal: killed
// exit status 1
