package composite

import (
	"fmt"
)

type Component interface {
	Search(keyword string)
}

type File struct {
	name string
}

func NewFile(name string) *File {
	return &File{name: name}
}

func (f *File) Search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

type Folder struct {
	components []Component
	name       string
}

func NewFolder(name string) *Folder {
	return &Folder{name: name}
}

func (f *Folder) Search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, c := range f.components {
		c.Search(keyword)
	}
}

func (f *Folder) Add(c Component) {
	f.components = append(f.components, c)
}
