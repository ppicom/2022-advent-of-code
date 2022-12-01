package app

import (
	"strconv"
	"strings"
)

func CalorieCounting(input string, nOfElves int) (int, error) {
	sanitizedInput := strings.TrimSuffix(input, "\n")
	sanitizedInput = strings.TrimPrefix(sanitizedInput, "\n")
	sanitizedInput = strings.ReplaceAll(sanitizedInput, "\t", "")
	listOfStringsReprOneElfSnacks := strings.Split(sanitizedInput, "\n\n")

	rankedCaloryTotals := make([]int, nOfElves)
	for _, stringReprOneElfSnacks := range listOfStringsReprOneElfSnacks {
		listOfStrReprOneSnack := strings.Split(stringReprOneElfSnacks, "\n")
		var caloriesCarriedByThisElf int
		for _, strReprOneSnack := range listOfStrReprOneSnack {
			caloriesOfASnack, err := strconv.ParseInt(strReprOneSnack, 10, 64)
			if err != nil {
				return 0, err
			}

			caloriesCarriedByThisElf += int(caloriesOfASnack)
		}

		for i := 0; i < len(rankedCaloryTotals); i++ {
			if caloriesCarriedByThisElf > rankedCaloryTotals[i] {
				for j := len(rankedCaloryTotals) - 1; j > i; j-- {
					rankedCaloryTotals[j] = rankedCaloryTotals[j-1]
				}

				rankedCaloryTotals[i] = caloriesCarriedByThisElf
				break
			}
		}
	}
	var grandTotal int
	for _, c := range rankedCaloryTotals {
		grandTotal += c
	}

	return grandTotal, nil
}
