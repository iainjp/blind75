package exercises

func maxArea(height []int) int {
	// two heights, only one solution; short-circuit
	if len(height) == 2 {
		return min(height[0], height[1]) // length is 1
	}

	maxArea := 0
	left := 0
	right := len(height) - 1

	// two pointer solution; calculate each candidate of left and right until you get the max, then return that
	for left < right {
		l := right - left
		// take min of two heights, as water will overflow, not slant
		h := min(height[left], height[right])
		maxArea = max(l*h, maxArea)

		// move pointer with lower height
		if height[left] < height[right] {
			left += 1
		} else {
			right -= 1
		}
	}

	return maxArea
}
