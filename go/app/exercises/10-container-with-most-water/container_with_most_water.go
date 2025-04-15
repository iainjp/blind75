package exercises

func maxArea(height []int) int {
	// two heights, only one solution
	if len(height) == 2 {
		return max(height[0], height[1]) // removes *1 which is the length of container
	}

	maxSize := 0
	// two pointer solution; calculate each candidate of left and right until you get the max, then return that
	for left := range height {
		right := len(height) - 1

		// until we've considered all candidates, calculate size, set if max, move right in 1
		for left < right {
			size := (right - left - 1) * height[right]
			maxSize = max(size, maxSize)
			right -= 1
		}
	}

	return maxSize
}
