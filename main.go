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

	score, err := app.RucksackReorganization(string(dat))
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	badgesPriority, err := app.BadgesPriority(string(dat))
	if err != nil {
		panic(err)
	}

	fmt.Printf("The first solution is: %d\n", score)
	fmt.Printf("The second solution is: %d\n", badgesPriority)
}
