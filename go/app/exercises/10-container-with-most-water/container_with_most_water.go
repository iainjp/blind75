package exercises

func maxArea(height []int) int {
	// two heights, only one solution
	if len(height) == 2 {
		return min(height[0], height[1]) // removes length as it's 1
	}

	maxSize := 0
	// two pointer solution; calculate each candidate of left and right until you get the max, then return that
	for left := range height {
		right := len(height) - 1

		maxHeightConsidered := min(left, right)

		// until we've considered all candidates, calculate size, set if max, move right in 1
		for left < right {
			// skip if we've already considered a heigher value
			if min(left, right) < maxHeightConsidered {
				right -= 1
				continue
			}

			maxHeightConsidered = min(left, right)

			l := right - left
			// take min of two heights, as water will overflow, not slant
			h := min(height[left], height[right])
			size := l * h
			maxSize = max(size, maxSize)
			right -= 1
		}
	}

	return maxSize
}
