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
		count, err := CalorieCounting(tt.input, 1)

		suite.NoError(err)
		suite.Equal(tt.calorieCount, count)
	}
}

func (suite *CaloriesTestSuite) Test_CalorieCounting_ThreeElves() {
	tests := []struct {
		input        string
		calorieCount int
	}{
		{
			input: `1

1

1`,
			calorieCount: 3,
		},
		{
			input: `1

2

1`,
			calorieCount: 4,
		},
		{
			input: `1

2
3

1`,
			calorieCount: 7,
		},
		{
			input: `1

2
3

1

10`,
			calorieCount: 16,
		},
		{
			input: `1

2
3

16

1

10`,
			calorieCount: 31,
		},
	}

	for _, tt := range tests {
		count, err := CalorieCounting(tt.input, 3)

		suite.NoError(err)
		suite.Equal(tt.calorieCount, count)
	}
}

func Test(t *testing.T) {
	suite.Run(t, new(CaloriesTestSuite))
}
