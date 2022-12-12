package entities

import (
	"github.com/thoas/go-funk"
)

type Directory struct {
	Size        int
	Files       []*File
	Directories []*Directory
	Name        string
	Parent      *Directory
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{Name: name, Size: 0, Parent: parent}
}

func (d *Directory) AddFile(file *File) {
	if len(d.Files) == 0 || d.Files == nil {
		d.Files = make([]*File, 1)
		d.Files[0] = file
	} else {
		d.Files = append(d.Files, file)
	}
	d.Size += file.Size
}
func (d *Directory) AddDirectory(dir *Directory) {
	if len(d.Directories) == 0 || d.Directories == nil {
		d.Directories = make([]*Directory, 1)
		d.Directories[0] = dir
	} else {
		d.Directories = append(d.Directories, dir)
	}

}

func (d Directory) FindDirectory(name string) *Directory {
	directory := funk.Find(d.Directories, func(dir *Directory) bool {
		return dir.Name == name
	})
	return directory.(*Directory)
}

func (d *Directory) CalculateSize() {
	size := 0
	for i := 0; i < len(d.Files); i++ {
		size += d.Files[i].Size
	}

	for i := 0; i < len(d.Directories); i++ {
		size += d.Directories[i].Size
	}

	d.Size = size
}
