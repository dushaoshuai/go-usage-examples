package main

import (
	"fmt"
	"log"
	"sync"
)

func main() {
	log.SetFlags(0)

	recoverMyself()
	fmt.Println()

	// A G cannot recover another G.
	// So this call will panic.
	recoverAnotherG()
}

func recoverMyself() {
	defer func() {
		if err := recover(); err != nil {
			log.Print(err)
		}
	}()

	panic("I don't want to panic.")

	// Output:
	// I don't want to panic.
}

// A G cannot recover another G.
func recoverAnotherG() {
	defer func() {
		if err := recover(); err != nil {
			log.Print("Got you.")
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		panic("I just want to panic!")
	}()

	wg.Wait()
}
