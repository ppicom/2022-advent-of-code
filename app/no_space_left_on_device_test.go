package app

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) TestSumOfSizes() {
	tests := []struct {
		input  string
		output int
	}{
		{
			input: `$ cd /
			$ ls
			10 file
			`,
			output: 10,
		},
		{
			input: `$ cd /
			$ ls
			10 file.two
			10 file.one
			`,
			output: 20,
		},
		{
			input: `$ cd /
			$ ls
			dir a
			10 file.two
			$ cd a
			10 file.one
			`,
			output: 30,
		},
		{
			input: `$ cd /
			$ ls
			dir a
			dir b
			10 file.two
			$ cd a
			$ ls
			10 file.one
			$ cd ..
			$ cd b
			$ ls
			10 file.three
			`,
			output: 50,
		},
		{
			input: `$ cd /
			$ ls
			dir a
			$ cd a
			$ ls
			dir b
			$ cd b
			$ ls
			1 file
			$ cd ..
			$ cd ..
			`,
			output: 3,
		},
		{
			input: `$ cd /
			$ ls
			dir btsgrbd
			1 vmpgqbcd
			$ cd btsgrbd
			$ ls
			dir cmfdm
			`,
			output: 1,
		},
		{
			input: `$ cd /
			$ ls
			dir a
			14848514 b.txt
			8504156 c.dat
			dir d
			$ cd a
			$ ls
			dir e
			29116 f
			2557 g
			62596 h.lst
			$ cd e
			$ ls
			584 i
			$ cd ..
			$ cd ..
			$ cd d
			$ ls
			4060174 j
			8033020 d.log
			5626152 d.ext
			7214296 k
			`,
			output: 95437,
		},
	}

	for _, tt := range tests {

		score := SumOfDirectoriesOfSize100000Max(tt.input)

		suite.Equal(tt.output, score)
	}
}

func TestShouldTest(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
