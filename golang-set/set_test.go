package set_test

import (
	"fmt"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

func Example_set() {
	languages := mapset.NewThreadUnsafeSet("Go", "C", "Java")

	for l := range languages.Iter() {
		fmt.Println(l)
	}

	// goroutine leak ?
	languages.Iter()
	time.Sleep(3 * time.Second)
	// Output:
}
