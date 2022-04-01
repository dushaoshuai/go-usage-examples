package fmt_test

import "fmt"

func ExampleSprintf() {
	msg := "New Year"
	msg = fmt.Sprintf(msg, "一年有 %d 天", 365)
	fmt.Println(msg)
	// Output:
	// New Year%!(EXTRA string=一年有 %d 天, int=365)
}
