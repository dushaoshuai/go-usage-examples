package main

import (
	"api-examples/golang/golang_path_filepath/routines"
	"fmt"
)

var paths = []string{
	".",
}

func main() {
	for _, path := range paths {
		fmt.Println(routines.Abs(path))
	}
}
