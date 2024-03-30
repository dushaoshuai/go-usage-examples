package file_operations

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

// LinkAll creates hard links to all targets in src.
// dst and src must be directories.
func LinkAll(dst, src string) error {
	isDir, err := IsDir(src)
	if err != nil {
		return err
	}
	if !isDir {
		return fmt.Errorf("helper: not a directory: %v", src)
	}

	isDir, err = IsDir(dst)
	if err != nil {
		return err
	}
	if !isDir {
		return fmt.Errorf("helper: not a directory: %v", dst)
	}

	err = filepath.WalkDir(src, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		dstPath := filepath.Join(dst, rel)

		if d.IsDir() {
			dirInfo, err := d.Info()
			if err != nil {
				return err
			}
			err = os.MkdirAll(dstPath, dirInfo.Mode().Perm())
			if err != nil {
				return err
			}
		} else {
			err = os.Link(path, dstPath)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("helper: LinkAll failed: %v", err)
	}

	return nil
}
