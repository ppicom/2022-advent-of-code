package app

type directory struct {
	name     string
	children []node
	parent   node
}

// SetParent implements node
func (d *directory) SetParent(parent node) {
	d.parent = parent
}

// Parent implements node
func (d *directory) Parent() node {
	return d.parent
}

// Append implements node
func (d *directory) Append(child node) error {
	d.children = append(d.children, child)
	return nil
}

func newDirectory(name string) node {
	return &directory{
		name:     name,
		children: make([]node, 0),
	}
}

func (d *directory) Children() []node {
	return d.children
}

func (d *directory) Name() string {
	return d.name
}

func (d *directory) Weight() (total int) {
	for _, child := range d.children {
		total += child.Weight()
	}

	return
}
