package app

import "strings"

func CharsToStartOfPacket(input string) (charsToPacket int) {

	signals := strings.Split(input, "")

	var found bool
	var i int
	for ; i < len(signals) && !found; i++ {

		lastFourObservedSignals := make(map[string]int, 0)
		for j := i; j < i+4; j++ {

			lastFourObservedSignals[signals[j]] += 1
		}

		found = len(lastFourObservedSignals) == 4
	}

	if found {
		charsToPacket = i + 3
	}

	return

}
