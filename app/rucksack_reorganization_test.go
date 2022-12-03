package app

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) TestRucksackReorganization() {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  `vJrwpWtwJgWrhcsFMMfFFhFp`,
			output: 16,
		},
		{
			input:  `jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL`,
			output: 38,
		},
		{
			input:  `PmmdzqPrVvPwwTWBwg`,
			output: 42,
		},
		{
			input:  `wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn`,
			output: 22,
		},
		{
			input:  `ttgJtRGJQctTZtZT`,
			output: 20,
		},
		{
			input:  `CrZsJsPPZsGzwwsLwLmpwMDw`,
			output: 19,
		},
		{
			input: `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`,
			output: 157,
		},
	}

	for _, tt := range tests {

		score, err := RucksackReorganization(tt.input)

		suite.NoError(err)
		suite.Equal(tt.output, score)
	}
}

func (suite *TestSuite) TestBadgesPriority() {
	tests := []struct {
		input  string
		output int
	}{
		{
			input: `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg`,
			output: 18,
		},
		{
			input: `wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`,
			output: 52,
		},
		{
			input: `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw`,
			output: 70,
		},
		{
			input: `LhfjhcdjcGdhFfdGfdjdvwCCZMvvLvWwMLCLSwZC
rDnsbmptPmlbQMCrQWQQBZQW
gltgVPngDPbptPsbPzVgmDldfTdfczThjJJjfMcJdFHjjH`,
			output: 39, // M
		},
	}

	for _, tt := range tests {

		score, err := BadgesPriority(tt.input)

		suite.NoError(err)
		suite.Equal(tt.output, score)
	}
}

func TestShouldTest(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
