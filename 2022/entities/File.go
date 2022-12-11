package entities

type File struct {
	Parent *Directory
	Name   string
	Size   int
}

func NewFile(parent *Directory, name string, size int) *File {
	return &File{Parent: parent, Name: name, Size: size}
}
