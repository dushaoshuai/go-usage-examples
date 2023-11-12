package main

import (
	"log"
)

// panic(nil) -> recover() -> nil
func main() {
	defer func() {
		if x := recover(); x != nil {
			log.Println(x)
		}
	}()

	panic(nil)
}
