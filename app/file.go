package app

import "fmt"

type file struct {
	weight int
	name   string
	parent node
}

// SetParent implements node
func (f *file) SetParent(parent node) {
	f.parent = parent
}

// Parent implements node
func (f *file) Parent() node {
	return f.parent
}

// Append implements node
func (f *file) Append(child node) error {
	return fmt.Errorf("cannot append to file")
}

func newFile(weight int, name string) node {
	return &file{
		weight: weight,
		name:   name,
	}
}

func (f *file) Weight() int {
	return f.weight
}

func (f *file) Name() string {
	return f.name
}

func (f *file) Children() []node {
	return nil
}
