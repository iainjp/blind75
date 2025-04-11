package exercises

func findMin(nums []int) int {
	// if nums is already sorted (e.g. nums[0] <= nums[-1]), return nums[0]
	if nums[0] <= nums[len(nums)-1] {
		return nums[0]
	}

	// adapted binary search: one half will not be sorted, that half contains min
	leftIndex := 0
	rightIndex := len(nums) - 1

	for leftIndex < rightIndex {
		middleIndex := (leftIndex + rightIndex) / 2

		// left half is sorted; min isn't there
		if nums[0] <= nums[middleIndex] {
			leftIndex = middleIndex + 1
		} else {
			// right half is not sorted, min is there
			rightIndex = middleIndex
		}
	}

	return nums[leftIndex]
}
