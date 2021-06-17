package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println("hello world!")

	cmd := exec.Command("/usr/bin/bash", "-c", "echo hello world!; echo hello world!")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Run()
}

// $ go run main1.go
// hello world!
// hello world!
// hello world!
