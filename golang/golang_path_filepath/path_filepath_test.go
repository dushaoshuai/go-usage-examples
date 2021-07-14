package golang_path_filepath_test

import (
	"fmt"
	"path/filepath"
	"testing"
)

var paths = []string{
	"./index.html",
}

func TestAbs(t *testing.T) {
	// Abs returns an absolute representation of path.
	fmt.Println(filepath.Abs(paths[0]))
}
