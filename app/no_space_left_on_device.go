package app

import (
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

	fs := buildFileSystemFromInput(input)

	size := &totalSize

	inorder(fs.Root(), func(n node) {
		w := n.Weight()
		if w < 100000 {
			*size += w
		}
	})

	return
}

func buildFileSystemFromInput(input string) fileSystem {
	input = strings.ReplaceAll(input, "\t", "")
	input = strings.TrimPrefix(input, "\n")
	input = strings.TrimSuffix(input, "\n")

	listOfLinesOfOutput := strings.Split(input, "\n")

	fs := runCommands(listOfLinesOfOutput)
	return fs
}

func runCommands(listOfLinesOfOutput []string) fileSystem {
	fs := newFileSystem()

	for _, lineOfOutput := range listOfLinesOfOutput {

		interpret(lineOfOutput, fs)
	}
	return fs
}

func interpret(lineOfOutput string, fs fileSystem) {
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

func inorder(n node, withNode func(node)) {

	if n == nil || n.Children() == nil || len(n.Children()) == 0 {
		return
	}

	nOfChildren := len(n.Children())

	for i := 0; i < nOfChildren-1; i++ {
		inorder(n.Children()[i], withNode)
	}

	withNode(n)

	inorder(n.Children()[nOfChildren-1], withNode)
}
