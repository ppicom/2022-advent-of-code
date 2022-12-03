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
	}{}

	for _, tt := range tests {

		score, err := BeingTested(tt.input)

		suite.NoError(err)
		suite.Equal(tt.output, score)
	}
}

func TestShouldTest(t *testing.T) {
	suite.Run(t, new(TestSuite))
}
