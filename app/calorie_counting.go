package app

import (
	"strconv"
	"strings"
)

func CalorieCounting(input string) (int, error) {
	sanitizedInput := strings.TrimSuffix(input, "\n")
	sanitizedInput = strings.TrimPrefix(sanitizedInput, "\n")
	sanitizedInput = strings.ReplaceAll(sanitizedInput, "\t", "")
	listOfStringsReprOneElfSnacks := strings.Split(sanitizedInput, "\n\n")

	var mostCaloriesCarriedByOneElf int
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

		if caloriesCarriedByThisElf > mostCaloriesCarriedByOneElf {
			mostCaloriesCarriedByOneElf = caloriesCarriedByThisElf
		}
	}

	return mostCaloriesCarriedByOneElf, nil
}
