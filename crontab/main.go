package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds())
	// Funcs are invoked in their own goroutine, asynchronously.
	_, err := c.AddFunc("*/2 * * * * *", func() { fmt.Println(time.Now(), "Every two seconds !!!") })
	if err != nil {
		panic(err)
	}
	_, err = c.AddFunc("*/3 * * * * *", func() { fmt.Println(time.Now(), "Every three seconds !!!") })
	if err != nil {
		panic(err)
	}
	_, err = c.AddFunc("*/5 * * * * *", func() {
		cmd := exec.Command("/usr/bin/bash", "-c", "date; echo Every five seconds !!!")
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Run()

	})
	if err != nil {
		panic(err)
	}
	c.Run()
}
