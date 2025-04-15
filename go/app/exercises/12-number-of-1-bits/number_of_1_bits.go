package exercises

import (
	"strconv"
)

func hammingWeight(n int) int {
	const oneValue = rune(49)

	bitSetCount := 0
	// convert int to string representation of binary form
	for _, c := range strconv.FormatInt(int64(n), 2) {
		// count the 1s
		if c == oneValue {
			bitSetCount += 1
		}
	}

	return bitSetCount
}
