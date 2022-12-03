package app

import (
	"fmt"
	"strings"
)

const (
	UPPERCASE_BASE int = 38
	LOWERCASE_BASE int = 96
)

func RucksackReorganization(input string) (int, error) {
	input = strings.TrimPrefix(input, "\n")
	input = strings.TrimSuffix(input, "\n")

	listOfRucksacks := strings.Split(input, "\n")
	var sumOfPriorities int
	for _, ruck := range listOfRucksacks {
		firstCompartment, secondCompartment := ruck[:len(ruck)/2], ruck[len(ruck)/2:]
		firstCompartmentList, secondCompartmentList := strings.Split(firstCompartment, ""), strings.Split(secondCompartment, "")

		var found string
		for _, itemInFirstComparment := range firstCompartmentList {

			for _, itemInSecondCompartment := range secondCompartmentList {

				if itemInFirstComparment == itemInSecondCompartment {
					found = itemInFirstComparment
				}
			}
		}

		switch {
		case found == "":
			// Do nothing
		default:
			sumOfPriorities += assignPriorityTo(found[0])
		}
	}

	return sumOfPriorities, nil
}

func BadgesPriority(input string) (int, error) {
	input = strings.TrimPrefix(input, "\n")
	input = strings.TrimSuffix(input, "\n")
	input = strings.ReplaceAll(input, "\t", "")

	var sumOfPriorities int

	listOfRucksacks := strings.Split(input, "\n")
	nOfGroups := len(listOfRucksacks) / 3

	for i := 0; i < nOfGroups; i++ {

		listOfRucksForThisGroup := listOfRucksacks[3*i : 3*i+3]

		firstRuck, secondRuck, thirdRuck := listOfRucksForThisGroup[0], listOfRucksForThisGroup[1], listOfRucksForThisGroup[2]

		common := findCommonItemsIn(firstRuck, secondRuck)
		badge := findCommonItemsIn(common, thirdRuck)

		switch {
		case badge == "":
			// Do nothing
		default:
			sumOfPriorities += assignPriorityTo(badge[0])
		}
	}

	return sumOfPriorities, nil
}

func findCommonItemsIn(firstRuck, secondRuck string) string {
	common := ""
	for _, itemInFirst := range strings.Split(firstRuck, "") {

		var seen bool = strings.Contains(common, itemInFirst)

		if !seen && strings.Contains(secondRuck, itemInFirst) {

			common = fmt.Sprintf("%s%s", common, itemInFirst)
		}
	}

	return common
}

func assignPriorityTo(item byte) int {
	switch {
	case item >= 'a':
		return int(item) - LOWERCASE_BASE
	default:
		return int(item) - UPPERCASE_BASE
	}
}
