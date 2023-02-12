package main

import "fmt"

func main() {
	var a string
	for {
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(a)
		}
	}
}
