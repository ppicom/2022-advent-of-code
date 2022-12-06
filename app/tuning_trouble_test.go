package app

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

func (suite *TestSuite) Test_FindStartOfPacket() {
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

// func (suite *TestSuite) Test_FindStartOfMessage() {
// 	tests := []struct {
// 		input  string
// 		output int
// 	}{
// 		{
// 			input:  "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
// 			output: 19,
// 		},
// 		{
// 			input:  "bvwbjplbgvbhsrlpgdmjqwftvncz",
// 			output: 23,
// 		},
// 		{
// 			input:  "nppdvjthqldpwncqszvftbrmjlhg",
// 			output: 23,
// 		},
// 		{
// 			input:  "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
// 			output: 29,
// 		},
// 		{
// 			input:  "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
// 			output: 26,
// 		},
// 	}

// 	for _, tt := range tests {

// 		score := CharsToStartOfMessage(tt.input)

// 		suite.Equal(tt.output, score)
// 	}
// }

func TestShouldTest(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
