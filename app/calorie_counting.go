package app

import (
	"strconv"
	"strings"
)

func CalorieCounting(input string) (int, error) {
	trimmedInput := strings.TrimSuffix(input, "\n")
	trimmedInput = strings.TrimPrefix(trimmedInput, "\n")
	trimmedInput = strings.ReplaceAll(trimmedInput, "\t", "")
	separated := strings.Split(trimmedInput, "\n\n")

	var biggestSum int
	for _, elveCalories := range separated {
		split := strings.Split(elveCalories, "\n")
		var sum int
		for _, intStr := range split {
			parsed, err := strconv.ParseInt(intStr, 10, 64)
			if err != nil {
				return 0, err
			}

			sum += int(parsed)
		}

		if sum > biggestSum {
			biggestSum = sum
		}
	}

	return biggestSum, nil
}
