package main

import (
	"fmt"
	"os"

	"github.com/ppicom/2022-advent-of-code/app"
)

func main() {
	dat, _ := os.ReadFile("./input.txt")

	markersToPacket := app.CharsToStartOfPacket(string(dat))

	fmt.Printf("There are %d markers to start of packet.\n", markersToPacket)
}
