package main

import (
	"errors"
	"flag"
	"io/fs"
	"log"
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
		os.Exit(1)
	}
}

type file struct {
	srcDir string
	dstDir string
	file   string
}

func main() {
	fileChan := make(chan file, 20)
	for i := 0; i < 48; i++ {
		go func() {
			defer func() {
				err := recover()
				if err != nil {
					log.Println(err)
				}
			}()

			for f := range fileChan {
				err := cpFile(f)
				if err != nil {
					log.Println(err)
				}
			}
		}()
	}

	err := checkDir()
	if err != nil {
		log.Println(err)
		return
	}

}

func checkDir() error {
	srcEntries, err := os.ReadDir(*src)
	if err != nil {
		return err
	}
	_ = srcEntries

	_, err = os.ReadDir(*dst)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}

	return nil
}

func cpFile(f file) error {

}
