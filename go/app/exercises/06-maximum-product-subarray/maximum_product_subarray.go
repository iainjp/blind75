package exercises

func maxProduct(nums []int) int {
	// Kadane's algorithm, adapted to handle min branch (for future -ve*-ve results)

	positiveProduct := nums[0]
	negativeProduct := nums[0]
	bestProduct := nums[0]

	for _, num := range nums[1:] {
		newPositiveProduct := max(num, positiveProduct*num, negativeProduct*num)
		negativeProduct = min(num, positiveProduct*num, negativeProduct*num)
		positiveProduct = newPositiveProduct

		bestProduct = max(newPositiveProduct, bestProduct)
	}

	return bestProduct
}
