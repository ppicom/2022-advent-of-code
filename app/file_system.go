package app

type fileSystem interface {
	Current() node
	Root() node
	MoveUp()
	MoveDownTo(subDirectory string)
	AddFile(weight int, name string)
	AddDirectory(name string)
}

type fs struct {
	current node
	root    node
}

// Root implements fileSystem
func (f *fs) Root() node {
	return f.root
}

// AddDirectory implements fileSystem
func (f *fs) AddDirectory(name string) {
	d := newDirectory(name)
	d.SetParent(f.Current())
	f.Current().Append(d)
}

// AddFile implements fileSystem
func (f *fs) AddFile(weight int, name string) {
	file := newFile(weight, name)
	file.SetParent(f.Current())
	f.Current().Append(file)
}

// Current implements fileSystem
func (f *fs) Current() node {
	return f.current
}

// MoveDownTo implements fileSystem
func (f *fs) MoveDownTo(subDirectory string) {
	for _, child := range f.current.Children() {
		if child.Name() == subDirectory {
			f.current = child
			return
		}
	}
}

// MoveUp implements fileSystem
func (f *fs) MoveUp() {
	f.current = f.current.Parent()
}

func newFileSystem() fileSystem {
	root := newDirectory("/")
	return &fs{
		current: root,
		root:    root,
	}
}
