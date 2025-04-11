package exercises

func maxSubArray(nums []int) int {
	// Kadane's algorithm

	bestSum := -100000
	currentSum := 0

	for _, num := range nums {
		currentSum = max(num, currentSum+num)
		bestSum = max(currentSum, bestSum)
	}

	return bestSum
}
