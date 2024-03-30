package file_operations

import (
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestLinkAll(t *testing.T) {
	src := ".."

	dst, err := os.MkdirTemp("", "*")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		os.RemoveAll(dst)
	})

	err = LinkAll(dst, src)
	if err != nil {
		t.Fatal(err)
	}

	walkDir := func(root string) ([]string, error) {
		var files []string

		err = filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				return err
			}
			rel, err := filepath.Rel(root, path)
			if err != nil {
				return err
			}
			files = append(files, rel)
			return nil
		})

		return files, err
	}

	srcFiles, err := walkDir(src)
	if err != nil {
		t.Fatal(err)
	}
	dstFiles, err := walkDir(dst)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(srcFiles, dstFiles) {
		t.Errorf("LinkAll() want = %v, got = %v", srcFiles, dstFiles)
	}
}
