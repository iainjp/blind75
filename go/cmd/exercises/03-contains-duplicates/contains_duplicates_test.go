package exercises

import (
	"testing"

	"github.com/iainjp/blind75/cmd/utils"
)

func TestContainsDuplicates(t *testing.T) {

	t.Run("ExampleOne", func(t *testing.T) {
		nums := []int{1, 2, 3, 1}
		want := true

		got := containsDuplicate(nums)

		utils.CheckEqual(got, want, t)
	})

	t.Run("ExampleTwo", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		want := false

		got := containsDuplicate(nums)

		utils.CheckEqual(got, want, t)
	})

	t.Run("ExampleThree", func(t *testing.T) {
		nums := []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}
		want := true

		got := containsDuplicate(nums)

		utils.CheckEqual(got, want, t)
	})
}
