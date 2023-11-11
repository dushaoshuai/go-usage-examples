package main

import (
	"errors"
	"flag"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"sync"
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
	fullPath    string
	perm        fs.FileMode
	dstFullPath string
}

func main() {
	err := checkDir()
	if err != nil {
		log.Println(err)
		return
	}

	var (
		fileChan = make(chan file, 20)
		wg       sync.WaitGroup
	)
	for i := 0; i < 48; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
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

	_ = filepath.WalkDir(*src,
		func(path string, d fs.DirEntry, err error) error {
			if err != nil {
				log.Println(err)
				if d.IsDir() {
					return fs.SkipDir
				}
				return nil
			}

			err = walkDirFunc(path, d, fileChan)
			if err != nil {
				log.Println(err)
				if d.IsDir() {
					return fs.SkipDir
				}
				return nil
			}

			return nil
		},
	)

	close(fileChan)
	wg.Wait()
	log.Println("cpdir: Done")
}

func checkDir() error {
	_, err := os.ReadDir(*src)
	if err != nil {
		return err
	}

	_, err = os.ReadDir(*dst)
	if err != nil && !errors.Is(err, fs.ErrNotExist) {
		return err
	}

	return nil
}

func cpFile(f file) error {
	data, err := os.ReadFile(f.fullPath)
	if err != nil {
		return err
	}
	return os.WriteFile(f.dstFullPath, data, f.perm)
}

func walkDirFunc(path string, d fs.DirEntry, fileChan chan<- file) error {
	rel, err := filepath.Rel(*src, path)
	if err != nil {
		return nil
	}
	dstFullPath := filepath.Join(*dst, rel)

	fileInfo, err := d.Info()
	if err != nil {
		return err
	}

	if d.IsDir() {
		err = os.MkdirAll(dstFullPath, fileInfo.Mode().Perm())
		if err != nil {
			return err
		}
	} else {
		fileChan <- file{
			fullPath:    path,
			perm:        fileInfo.Mode().Perm(),
			dstFullPath: dstFullPath,
		}
	}

	return nil
}
