package sha1_test

import (
	"crypto/sha1"
	"fmt"
)

func ExampleSha1Sum() {
	data := []byte("This page intentionally left blank.")
	fmt.Printf("% x", sha1.Sum(data))
	// Output:
	// af 06 49 23 bb f2 30 15 96 aa c4 c2 73 ba 32 17 8e bc 4a 96
}
