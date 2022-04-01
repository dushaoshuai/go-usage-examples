package fmt_test

import "fmt"

func Example_width_precision() {
	// fmt.Printf("%3.-2f\n", 4.5) ==> Printf format %3.- has unknown verb -
	fmt.Printf("%9v\n", 8)
	fmt.Printf("%9.3v\n", "gophers")
	fmt.Printf("%9.3x\n", "gophers")
	fmt.Printf("%9.3X\n", "gophers")
	fmt.Printf("%9.3v\n", 5.6)
	fmt.Printf("%9.3f\n", 5.6)
	fmt.Printf("%6.3g\n", 12.345)
	// Output:
	//         8
	//       gop
	//    676f70
	//    676F70
	//       5.6
	//     5.600
	//   12.3
}
