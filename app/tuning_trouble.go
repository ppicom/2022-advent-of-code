package app

import "strings"

const (
	START_OF_PACKET_MARKER_LEN int = 4
	START_OF_MSG_MARKER_LEN    int = 14
)

func CharsToStartOfPacket(input string) (charsToPacket int) {

	return findMarker(START_OF_PACKET_MARKER_LEN, input)
}

func CharsToStartOfMessage(input string) (charsToMsg int) {

	return findMarker(START_OF_MSG_MARKER_LEN, input)
}

func findMarker(marker int, input string) (charsToPacket int) {

	signals := strings.Split(input, "")

	var (
		found bool
		where int
	)

	for i := 0; i < len(signals) && !found; i++ {

		found, where = markerStartsAt(i, marker, signals)
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
