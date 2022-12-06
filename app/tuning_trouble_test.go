package app

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) TestPlay() {
	tests := []struct {
		input  string
		output int
	}{
		{
			input:  "bvwbjplbgvbhsrlpgdmjqwftvncz",
			output: 5,
		},
		{
			input:  "nppdvjthqldpwncqszvftbrmjlhg",
			output: 6,
		},
		{
			input:  "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			output: 10,
		},
		{
			input:  "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			output: 11,
		},
	}

	for _, tt := range tests {

		score := CharsToStartOfPacket(tt.input)

		suite.Equal(tt.output, score)
	}
}

func TestShouldTest(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
