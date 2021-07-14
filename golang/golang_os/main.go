package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	// the current directory is where you execute the program,
	// not where the program resides.
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dir)

	// os.ExpandEnv()
	fmt.Println(os.ExpandEnv("${HOME}"))
}
