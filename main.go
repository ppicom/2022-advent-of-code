package main

import (
	"fmt"
	"os"

	"github.com/ppicom/2022-advent-of-code/app"
)

func main() {
	dat, _ := os.ReadFile("./input.txt")

	signalsToPacket := app.CharsToStartOfPacket(string(dat))

	fmt.Printf("There are %d markers to start of packet.\n", signalsToPacket)

	signalsToMsg := app.CharsToStartOfMessage(string(dat))

	fmt.Printf("There are %d markers to start of message.\n", signalsToMsg)
}
