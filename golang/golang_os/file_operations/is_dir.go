package file_operations

import (
	"os"
)

// IsDir reports whether name is a directory.
func IsDir(dirName string) (bool, error) {
	info, err := os.Stat(dirName)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}
