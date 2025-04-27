package exercises

import "slices"

func insert(intervals [][]int, newInterval []int) [][]int {
	const LOWER, UPPER = 0, 1

	// result := [][]int{}

	var insertInterval = func(intervals [][]int, newInterval []int) [][]int {
		result := [][]int{}

		for _, existingInterval := range intervals {
			// determine if we have an overlap -- yeah this is bogus.
			if newInterval[LOWER] <= existingInterval[LOWER] || newInterval[UPPER] >= existingInterval[UPPER] {
				// we have an overlap; merge it
				combined := append(newInterval, existingInterval...)
				newLower := slices.Min(combined)
				newUpper := slices.Max(combined)
				result = append(result, []int{newLower, newUpper})
			} else {
				// otherwise, add existing interval
				result = append(result, existingInterval)
			}
		}

		return result
	}

	// I think this requires multiple iterations, to resolve new overlaps as they get generated ...

	inserted := insertInterval(intervals, newInterval)

	return inserted
}
