package file_operations

import (
	"fmt"
	"io"
	"os"
)

// CpFile copy a single file.
func CpFile(dst, src string) error {
	s, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("helper: %v", err)
	}
	defer s.Close()

	d, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("helper: %v", err)
	}
	defer d.Close()

	_, err = io.Copy(d, s)
	if err != nil {
		return fmt.Errorf("helper: cannot copy file: %v", err)
	}
	return nil
}
