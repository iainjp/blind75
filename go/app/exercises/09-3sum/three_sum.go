package exercises

import "slices"

func threeSum(nums []int) [][]int {
	// similar to twoSum, but triplets; we need to fix one of the numbers
	// can use -nums[i], then goal becomes nums[j]+nums[k]==-nums[i], such that
	// sum of all three is 0

	results := [][]int{}

	slices.Sort(nums)

	for curr := range nums {
		target := -nums[curr]

		left := curr + 1
		right := len(nums) - 1

		for left < right {
			// skip to avoid duplicates - this is tricky.
			if curr > 0 && nums[curr] == nums[curr-1] {
				curr += 1
				continue
			} else if nums[left] == nums[left-1] {
				left += 1
				continue
			} else if right < len(nums)-1 && nums[right] == nums[right+1] {
				right -= 1
				continue
			}

			sum := nums[left] + nums[right]

			switch {
			// we find a valid triplet
			case sum == target:
				results = append(results, []int{nums[curr], nums[left], nums[right]})
				left += 1
				right -= 1
			// too low; we need to move j
			case sum < target:
				left += 1
			// too high, we need to move k
			case sum > target:
				right -= 1
			}

		}
	}

	return results
}
