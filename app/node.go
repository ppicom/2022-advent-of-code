package app

type node interface {
	Weight() int
	Name() string
	Children() []node
	Parent() node
	SetParent(parent node)
	Append(child node) error
}
