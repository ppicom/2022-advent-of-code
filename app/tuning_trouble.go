package app

import "strings"

const (
	START_OF_PACKET_MARKER_LEN int = 4
)

func CharsToStartOfPacket(input string) (charsToPacket int) {

	signals := strings.Split(input, "")

	var (
		found bool
		where int
	)

	for i := 0; i < len(signals) && !found; i++ {

		found, where = markerStartsAt(i, START_OF_PACKET_MARKER_LEN, signals)
	}

	if found {
		charsToPacket = where
	}

	return

}

func markerStartsAt(position, markerLen int, signals []string) (found bool, where int) {

	lastObservedSignals := make(map[string]int, 0)
	for j := position; j < position+markerLen; j++ {

		lastObservedSignals[signals[j]] += 1
	}

	found = len(lastObservedSignals) == markerLen

	return found, position + markerLen
}
