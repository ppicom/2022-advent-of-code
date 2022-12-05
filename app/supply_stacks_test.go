package app

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) TestCreateStacks() {
	tests := []struct {
		input  string
		output string
	}{
		{
			input: `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

`,
			output: "NDP",
		},
		{
			input: `    [D]    
[P] [C]    
[Z] [M] [N]
 1   2   3 

`,
			output: "PDN",
		},
		{
			input: `            [J] [Z] [G]            
            [Z] [T] [S] [P] [R]    
[R]         [Q] [V] [B] [G] [J]    
[W] [W]     [N] [L] [V] [W] [C]    
[F] [Q]     [T] [G] [C] [T] [T] [W]
[H] [D] [W] [W] [H] [T] [R] [M] [B]
[T] [G] [T] [R] [B] [P] [B] [G] [G]
[S] [S] [B] [D] [F] [L] [Z] [N] [L]
 1   2   3   4   5   6   7   8   9 

`,
			output: "RWWJZGPRW",
		},
	}

	for _, tt := range tests {

		topmostCases := MoveSupplyStacks(tt.input)

		suite.Equal(tt.output, topmostCases)
	}
}

func (suite *TestSuite) TestMoveCratesBetweenStacks() {
	tests := []struct {
		input  string
		output string
	}{
		{
			input: `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`,
			output: "CMZ",
		},
	}

	for _, tt := range tests {

		topmostCases := MoveSupplyStacks(tt.input)

		suite.Equal(tt.output, topmostCases)
	}
}

func TestShouldTest(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
