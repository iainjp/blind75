package exercises

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool, len(nums))

	for _, num := range nums {
		if _, found := set[num]; found {
			return true
		}
		set[num] = true
	}

	return false
}
