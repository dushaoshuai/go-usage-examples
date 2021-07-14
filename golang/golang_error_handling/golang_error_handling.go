package golang_error_handling

import (
	"fmt"
	"log"
	"os"
	"runtime"
)

var skip = 100
var skipStart = 1
var exitStatusCode = 1

func LogError(e error) {
	if e == nil {
		return
	}
	log.Print(e)
	// https://github.com/gohouse/e/blob/3cf103d33981f87b78674612d219538c5ea4f624/context.go#L45
	for i := skipStart; i < skip; i++ {
		if function, file, line, ok := runtime.Caller(i); ok {
			fmt.Println(runtime.FuncForPC(function).Name())
			fmt.Printf("    %s:%d\n", file, line)
		} else {
			break
		}
	}
	os.Exit(exitStatusCode)
}
