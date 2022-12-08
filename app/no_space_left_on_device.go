package app

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	is_list_cmd  *regexp.Regexp = regexp.MustCompile(`^\$ ls`)
	is_cd_moveup *regexp.Regexp = regexp.MustCompile(`^\$ cd \.\.`)
	is_cd_root   *regexp.Regexp = regexp.MustCompile(`^\$ cd /`)
	is_cd_else   *regexp.Regexp = regexp.MustCompile(`^\$ cd `)
	is_dir       *regexp.Regexp = regexp.MustCompile(`^dir `)
	is_file      *regexp.Regexp = regexp.MustCompile(`^\d* `)
)

func SumOfDirectoriesOfSize100000Max(input string) (totalSize int) {

	input = strings.ReplaceAll(input, "\t", "")
	input = strings.TrimPrefix(input, "\n")
	input = strings.TrimSuffix(input, "\n")

	listOfLinesOfOutput := strings.Split(input, "\n")

	fs := newFileSystem()

	for _, lineOfOutput := range listOfLinesOfOutput {

		switch {
		case is_cd_root.MatchString(lineOfOutput):
			// Do nothing
		case is_list_cmd.MatchString(lineOfOutput):
			//DO nothing
		case is_cd_moveup.MatchString(lineOfOutput):
			fs.MoveUp()
		case is_cd_else.MatchString(lineOfOutput):
			cmdParts := strings.Split(lineOfOutput, " ")
			dirName := cmdParts[len(cmdParts)-1]
			fs.MoveDownTo(dirName)
		case is_dir.MatchString(lineOfOutput):
			cmdParts := strings.Split(lineOfOutput, " ")
			dirName := cmdParts[len(cmdParts)-1]
			fs.AddDirectory(dirName)
		case is_file.MatchString(lineOfOutput):
			cmdParts := strings.Split(lineOfOutput, " ")
			wghtStr, name := cmdParts[0], cmdParts[1]
			wght, _ := strconv.ParseInt(wghtStr, 10, 64)
			fs.AddFile(int(wght), name)
		}

	}

	size := &totalSize
	inorder(fs.Root(), size)

	return
}

type node interface {
	Weight() int
	Name() string
	Children() []node
	Parent() node
	SetParent(parent node)
	Append(child node) error
}

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

func inorder(n node, size *int) {

	if n == nil || n.Children() == nil || len(n.Children()) == 0 {
		return
	}

	nOfChildren := len(n.Children())

	for i := 0; i < nOfChildren-1; i++ {
		inorder(n.Children()[i], size)
	}

	w := n.Weight()
	if w < 100000 {
		*size += w
	}

	inorder(n.Children()[nOfChildren-1], size)
}
