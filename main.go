package main

import (
	"fmt"
	"os"

	"github.com/ppicom/2022-advent-of-code/app"
)

func main() {
	dat, err := os.ReadFile("./input.txt")
	if err != nil {
		panic(err)
	}

	size := app.SumOfDirectoriesOfSize100000Max(string(dat))

	fmt.Printf("The total size of directories below 100000 is %d\n", size)
}
