package exercises

func search(nums []int, target int) int {
	// adapted binary search; one half will not be sorted
	leftIndex := 0
	rightIndex := len(nums) - 1

	for leftIndex <= rightIndex {
		middleIndex := (leftIndex + rightIndex) / 2

		// check for target
		if nums[middleIndex] == target {
			return middleIndex
		}

		// consider which half to discard, based on whether it's a sorted half AND target in that range
		// if left half sorted
		if nums[leftIndex] <= nums[middleIndex] {
			// if in range, discard right
			if nums[leftIndex] <= target && target < nums[middleIndex] {
				rightIndex = middleIndex - 1
			} else {
				// not in sorted range; discard left
				leftIndex = middleIndex + 1
			}
		} else {
			// if in range, discard left
			if nums[middleIndex] < target && target <= nums[rightIndex] {
				leftIndex = middleIndex + 1
			} else {
				// if not in sorted range, discard right
				rightIndex = middleIndex - 1
			}
		}
	}

	return -1
}
