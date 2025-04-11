package exercises

func productExceptSelf(nums []int) []int {
	// track product per index
	products := make([]int, len(nums))

	// keeping track of the accumulated product
	leftProduct := 1
	rightProduct := 1

	// starting left, track running product and assign to `products[i]`,
	for i := 0; i < len(nums); i++ {
		products[i] = leftProduct
		leftProduct *= nums[i]
	}

	// start right, track running product and multiply and assign to `products[i]`
	for i := len(nums) - 1; i >= 0; i-- {
		products[i] = products[i] * rightProduct
		rightProduct *= nums[i]
	}

	return products
}
