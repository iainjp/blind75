package exercises

func productExceptSelf(nums []int) []int {
	// track product per index
	products := make([]int, len(nums))
	for i := range products {
		products[i] = 1
	}

	for i, val := range nums {
		for otherI := range len(nums) {
			if otherI == i {
				continue
			}

			products[otherI] = val * products[otherI]
		}
	}

	return products
}
