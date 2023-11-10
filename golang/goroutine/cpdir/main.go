package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
)

var (
	src = flag.String("src", "", "source directory")
	dst = flag.String("dst", "", "destination directory")
)

func init() {
	flag.Parse()
	if *src == "" || *dst == "" {
		flag.PrintDefaults()
	}
}

func main() {

	for i := 0; i < 48; i++ {

	}
}

func checkDir() error {
	srcDir, err := os.Open(*src)
	if err != nil {
		return err
	}
	srcDirInfo, err := srcDir.Stat()
	if err != nil {
		return err
	}
	if !srcDirInfo.IsDir() {
		return fmt.Errorf("cpdir: src %s is not a directory", srcDirInfo.Name())
	}

	os.ReadDir()

	dstDirInfo, err := os.Stat(*dst)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			return nil
		}
		return err
	}
	if !dstDirInfo.IsDir() {
		return fmt.Errorf("cpdir: dst %s is not a directory", dstDirInfo.Name())
	}

}
