package app

import (
	"fmt"
	"strings"
)

func RockPaperScissorsStrategy(input string) (int, error) {
	sanitizedInput := strings.TrimPrefix(input, "\n")
	sanitizedInput = strings.TrimSuffix(sanitizedInput, "\n")
	sanitizedInput = strings.ReplaceAll(sanitizedInput, "\t", "")
	listOfGames := strings.Split(sanitizedInput, "\n")

	var myScoreForAllGames int
	for _, aGame := range listOfGames {
		strategyOfAGame := strings.Split(aGame, " ")

		opponentsHand, mustEndIn := strategyOfAGame[0], strategyOfAGame[1]

		var myScoreForThisHand int
		var myHand string
		switch mustEndIn {
		case "X":
			myHand = LoseTo(opponentsHand)
		case "Z":
			myHand = WinTo(opponentsHand)
		case "Y":
			myHand = Draw(opponentsHand)
		}

		play := fmt.Sprintf("%s %s", opponentsHand, myHand)
		score, err := RockPaperScissors(play)
		if err != nil {
			return 0, err
		}

		myScoreForThisHand = score

		myScoreForAllGames += myScoreForThisHand

	}

	return myScoreForAllGames, nil

}

func LoseTo(otherHand string) string {
	switch otherHand {
	case "A":
		return "Z"
	case "B":
		return "X"
	case "C":
		fallthrough
	default:
		return "Y"
	}
}

func Draw(otherHand string) string {
	switch otherHand {
	case "A":
		return "X"
	case "B":
		return "Y"
	case "C":
		fallthrough
	default:
		return "Z"
	}
}

func WinTo(otherHand string) string {
	switch otherHand {
	case "A":
		return "Y"
	case "B":
		return "Z"
	case "C":
		fallthrough
	default:
		return "X"
	}
}

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
