package routines

import "path/filepath"

func Abs(path string) (string, error) {
	// The result depends on where you execute command
	// such as "go run main.go", not where main.go resides.
	return filepath.Abs(path)
}
