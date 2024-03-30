package file_operations

import (
	"io/fs"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestIsDir(t *testing.T) {
	tests := []struct {
		name    string
		dirName string
		want    bool
		wantErr bool
	}{
		{"file", "./link_test.go", false, false},
		{"dir", ".", true, false},
		{"non-exist", "/tmp/non-exist", false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := IsDir(tt.dirName)
			if (err != nil) != tt.wantErr {
				t.Errorf("IsDir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("IsDir() got = %v, want %v", got, tt.want)
			}
		})
	}
}

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
