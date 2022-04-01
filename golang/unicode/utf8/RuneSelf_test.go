package utf8

import (
	"fmt"
	"unicode/utf8"
)

// utf8.RuneSelf == 0x80 == 128
// runes 0-127          用 1 个字节表示
// runes 128-2047       用 2 个字节表示
// runes 2048-65535     用 3 个字节表示
// runes 65536-0x10ffff 用 4 个字节表示
// utf8.MaxRune == '\U0010FFFF'
// uft8.UTFMax == 4
// utf8.RuneError == '\uFFFD'

func ExampleRuneSelf() {
	fmt.Printf("%c\n", utf8.RuneSelf+1)
	fmt.Printf("%c\n", utf8.RuneSelf)
	fmt.Printf("%c\n", utf8.RuneSelf-2)
	// Output:
	// 
	// 
	// ~
}
