package main

import (
	"fmt"
	"os"

	"github.com/ppicom/2022-advent-of-code/app"
)

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	score, err := app.RockPaperScissors(string(dat))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	secondScore, err := app.RockPaperScissorsStrategy(string(dat))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Printf("Your score is %d\n", score)
	fmt.Printf("Your score with the second strategy is %d\n", secondScore)
}
