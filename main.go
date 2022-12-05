package main

import (
	"fmt"
	"os"

	"github.com/ppicom/2022-advent-of-code/app"
)

func main() {
	data, err := os.ReadFile(("./input.txt"))
	if err != nil {
		panic(err)
	}

	crates := app.MoveSupplyStacks(string(data))

	fmt.Printf("The topmost crates are: %s\n", crates)

	crates = app.MoveSupplyStacksUsingNewCrane(string(data))

	fmt.Printf("The topmost crates now are: %s\n", crates)
}
