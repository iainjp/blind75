package exercises

import (
	"testing"

	"github.com/iainjp/blind75/cmd/utils"
)

func TestTwoSums(t *testing.T) {

	t.Run("ExampleOne", func(t *testing.T) {
		nums := []int{2, 7, 11, 15}
		target := 9
		want := []int{0, 1}

		got := twoSum(nums, target)

		utils.CheckEqual(got, want, t)
	})

	t.Run("ExampleTwo", func(t *testing.T) {
		nums := []int{3, 2, 4}
		target := 6
		want := []int{1, 2}

		got := twoSum(nums, target)

		utils.CheckEqual(got, want, t)
	})

	t.Run("ExampleThree", func(t *testing.T) {
		nums := []int{3, 3}
		target := 6
		want := []int{0, 1}

		got := twoSum(nums, target)

		utils.CheckEqual(got, want, t)
	})
}
