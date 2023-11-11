package files_read_write_test

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
)

func Example_filepath_WalkDir() {
	filepath.WalkDir("..", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Println(err)
			if d.IsDir() {
				return fs.SkipDir
			}
			return nil
		}

		absPath, _ := filepath.Abs(path)
		fmt.Println(absPath)
		return nil
	})

	// Output:
}
