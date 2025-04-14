package exercises

func search(nums []int, target int) int {
	// adapted binary search; one half will not be sorted
	leftIndex := 0
	rightIndex := len(nums) - 1
	middleIndex := 0

	for leftIndex < rightIndex {
		middleIndex = (leftIndex + rightIndex) / 2

		// consider which half to discard, based on whether it's a sorted half AND target in that range
		// if left half sorted
		if nums[0] <= nums[middleIndex] {
			// if in range, discard right
			if nums[0] <= target && target <= nums[middleIndex] {
				rightIndex = middleIndex
			} else {
				// not in sorted range; discard left
				leftIndex = middleIndex + 1
			}
		}

		// if right half sorted
		if nums[middleIndex] <= nums[rightIndex] {
			// if in range, discard left
			if nums[middleIndex] <= target && target <= nums[rightIndex] {
				leftIndex = middleIndex + 1
			} else {
				// if not in sorted range, discard right
				rightIndex = middleIndex
			}
		}
	}

	if nums[middleIndex] == target {
		return middleIndex
	} else {
		return -1
	}
}
