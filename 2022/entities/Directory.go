package entities

type Directory struct {
	Size   int
	Files  []*File
	Name   string
	Parent *Directory
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
