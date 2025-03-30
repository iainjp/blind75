package exercises

import "fmt"

type Pair struct {
	first  int
	second int
}

func getIndexPermutationPairs(nums []int) []Pair {
	// get all pairs of possible indices
	indexPairs := []Pair{}

	// get all pairs of possible indices
	for i := range len(nums) {

		for j := i + 1; j < len(nums); j++ {
			pair := Pair{first: i, second: j}
			indexPairs = append(indexPairs, pair)
		}
	}

	fmt.Println(indexPairs)

	return indexPairs
}

func twoSum(nums []int, target int) []int {
	var results []int
	indexPairs := getIndexPermutationPairs(nums)

	for _, pair := range indexPairs {
		if nums[pair.first]+nums[pair.second] == target {
			results = []int{pair.first, pair.second}
		}
	}

	return results
}
