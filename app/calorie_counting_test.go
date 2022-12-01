package app

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CaloriesTestSuite struct {
	suite.Suite
}

func (suite *CaloriesTestSuite) Test_CalorieCounting() {
	tests := []struct {
		input        string
		calorieCount int
	}{
		{
			input:        `1000`,
			calorieCount: 1000,
		},
		{
			input:        `2000`,
			calorieCount: 2000,
		},
		{
			input: `1000
2000`,
			calorieCount: 3000,
		},
		{
			input: `1000
2000
3000`,
			calorieCount: 6000,
		},
		{
			input: `1000
2000
3000

100`,
			calorieCount: 6000,
		},
		{
			input: `1000
2000
3000

10000`,
			calorieCount: 10000,
		},
		{
			input: `1000
2000
3000

10000

10000
20000`,
			calorieCount: 30000,
		},
		{
			input: `1000
2000
3000

10000

10000
20000
`,
			calorieCount: 30000,
		},
		{
			input: `
1000
2000
3000

10000

10000
20000
`,
			calorieCount: 30000,
		},
		{
			input: `
1000
	2000
3000

10000

10000
20000
`,
			calorieCount: 30000,
		},
	}

	for _, tt := range tests {
		count, err := CalorieCounting(tt.input)

		suite.NoError(err)
		suite.Equal(tt.calorieCount, count)
	}
}

func Test(t *testing.T) {
	suite.Run(t, new(CaloriesTestSuite))
}
