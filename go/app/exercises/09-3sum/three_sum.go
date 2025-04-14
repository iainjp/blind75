package exercises

import (
	"maps"
	"slices"
)

func threeSum(nums []int) [][]int {
	// similar to twoSum, but triplets; we need to fix one of the numbers
	// can use -nums[i], then goal becomes nums[j]+nums[k]==-nums[i], such that
	// sum of all three is 0

	// track what we've seen; definitely the biggest performance hit
	type Triplet struct {
		curr  int
		left  int
		right int
	}
	resultMap := make(map[Triplet]bool)
	slices.Sort(nums)

	for curr := range len(nums) - 1 {
		target := -nums[curr]

		left := curr + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[left] + nums[right]

			switch {
			// we find a valid triplet
			case sum == target:
				resultMap[Triplet{nums[curr], nums[left], nums[right]}] = true
				left += 1
				right -= 1
			// too low; we need to move left pointer
			case sum < target:
				left += 1
			// too high, we need to move right pointer
			case sum > target:
				right -= 1
			}
		}
	}

	results := [][]int{}
	for t := range maps.Keys(resultMap) {
		results = append(results, []int{t.curr, t.left, t.right})
	}
	return results
}
