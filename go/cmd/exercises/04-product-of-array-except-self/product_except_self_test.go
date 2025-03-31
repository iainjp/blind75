package exercises

import (
	"testing"

	"github.com/iainjp/blind75/cmd/utils"
)

func TestProductExceptSelf(t *testing.T) {
	t.Run("ExampleOne", func(t *testing.T) {
		nums := []int{1, 2, 3, 4}
		want := []int{24, 12, 8, 6}

		got := productExceptSelf(nums)

		utils.CheckEqual(got, want, t)
	})

	t.Run("ExampleTwo", func(t *testing.T) {
		nums := []int{-1, 1, 0, -3, 3}
		want := []int{0, 0, 9, 0, 0}

		got := productExceptSelf(nums)

		utils.CheckEqual(got, want, t)
	})
}
