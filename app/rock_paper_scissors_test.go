package app

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type RockPaperScissorsTestSuite struct {
	suite.Suite
}

func (suite *RockPaperScissorsTestSuite) TestPlay() {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  `A Y`,
			output: 8,
		},
		{
			input:  `B X`,
			output: 1,
		},
		{
			input: `A Y
B X`,
			output: 9,
		},
		{
			input: `A Y
B X
C Z`,
			output: 15,
		},
		{
			input:  `A X`,
			output: 4,
		},
		{
			input:  `B Y`,
			output: 5,
		},
		{
			input:  `C Z`,
			output: 6,
		},
		{
			input:  `A Z`,
			output: 3,
		},
		{
			input:  `B X`,
			output: 1,
		},
		{
			input:  `C Y`,
			output: 2,
		},
		{
			input:  `C X`,
			output: 7,
		},
		{
			input:  `B Z`,
			output: 9,
		},
		{
			input: `B Z
`,
			output: 9,
		},
	}

	for _, tt := range tests {

		score, err := RockPaperScissors(tt.input)

		suite.NoError(err)
		suite.Equal(tt.output, score)
	}
}

func (suite *RockPaperScissorsTestSuite) TestPlaySecondStrategy() {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  `A Y`,
			output: 4,
		},
		{
			input:  `B X`,
			output: 1,
		},
		{
			input:  `C Z`,
			output: 7,
		},
		{
			input: `A Y
B X
C Z`,
			output: 12,
		},
	}

	for _, tt := range tests {

		score, err := RockPaperScissorsStrategy(tt.input)

		suite.NoError(err)
		suite.Equal(tt.output, score)
	}
}

func TestShouldTest(t *testing.T) {
	suite.Run(t, new(RockPaperScissorsTestSuite))
}
