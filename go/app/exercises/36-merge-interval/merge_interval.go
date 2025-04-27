package exercises

import (
	"slices"
)

func merge(intervals [][]int) [][]int {
	const LOWER, UPPER = 0, 1

	// sort by first element in intervals
	var firstElemCmp = func(firstInterval []int, secondInterval []int) int {
		return firstInterval[0] - secondInterval[0]
	}
	slices.SortFunc(intervals, firstElemCmp)

	merged := [][]int{intervals[0]}

	for _, interval := range intervals[1:] {
		lastInserted := merged[len(merged)-1]

		// no overlap, just add it
		if lastInserted[UPPER] < interval[LOWER] {
			merged = append(merged, interval)
		} else {
			// overlapping intervals, possibly extend last inserted by taking max of both uppers
			lastInserted[UPPER] = max(lastInserted[UPPER], interval[UPPER])
		}
	}

	return merged
}
