package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	// the current directory is where you execute the program,
	// not where the program resides.
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(wd)

	// os.ExpandEnv()
	fmt.Println(os.ExpandEnv("${HOME}"))
	// Output:
	// /home/user_name

	// fetch the path name for the executable that started the current process.
	// The main use case is finding resources located relative to an executable.
	// go run main.go will get a path started with "/tmp".
	// Better go build then run.
	path, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	path, err = filepath.EvalSymlinks(path)
	if err != nil {
		log.Fatal(err)
	}
	dir := filepath.Dir(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("the path name for the executable that started the current process is : " + dir)
}
