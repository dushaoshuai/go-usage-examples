package file_operations

import (
	"io/fs"
	"os"
	"path/filepath"
)

// CpAll copy files from src to dst.
// dst and src must be directories.
func CpAll(dst, src string) error {
	isDir, err := IsDir(src)
	if err != nil {
		return err
	}
	if !isDir {
		return err
	}

	isDir, err = IsDir(dst)
	if err != nil {
		return err
	}
	if !isDir {
		return err
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
			err = CpFile(dstPath, path)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
