package main

import (
	"fmt"

	"github.com/dushaoshuai/go-usage-examples/golang/golang_path_filepath/routines"
)

var paths = []string{
	".",
}

func main() {
	for _, path := range paths {
		fmt.Println(routines.Abs(path))
	}
}
