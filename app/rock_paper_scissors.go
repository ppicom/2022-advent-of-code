package app

import "strings"

func RockPaperScissors(input string) (int, error) {
	sanitizedInput := strings.TrimPrefix(input, "\n")
	sanitizedInput = strings.TrimSuffix(sanitizedInput, "\n")
	sanitizedInput = strings.ReplaceAll(sanitizedInput, "\t", "")
	differentHands := strings.Split(sanitizedInput, "\n")

	var myScoreForAllHands int
	for _, hand := range differentHands {
		split := strings.Split(hand, " ")

		opponentsHand, myHand := split[0], split[1]

		var myScoreForThisHand int

		switch myHand {
		case "X":
			myScoreForThisHand += 1

			if opponentsHand == "A" {
				myScoreForThisHand += 3
			} else if opponentsHand == "C" {
				myScoreForThisHand += 6
			}
		case "Y":
			myScoreForThisHand += 2

			if opponentsHand == "B" {
				myScoreForThisHand += 3
			} else if opponentsHand == "A" {
				myScoreForThisHand += 6
			}
		case "Z":
			myScoreForThisHand += 3

			if opponentsHand == "C" {
				myScoreForThisHand += 3
			} else if opponentsHand == "B" {
				myScoreForThisHand += 6
			}
		}

		myScoreForAllHands += myScoreForThisHand

	}

	return myScoreForAllHands, nil

}
