package exercises

import (
	"testing"

	"github.com/iainjp/blind75/cmd/utils"
)

func TestTimeToBuyStocks(t *testing.T) {
	t.Run("ExampleOne", func(t *testing.T) {
		prices := []int{7, 1, 5, 3, 6, 4}
		want := 5

		got := maxProfit(prices)

		utils.CheckEqual(got, want, t)
	})

	t.Run("ExampleTwo", func(t *testing.T) {
		prices := []int{7, 6, 4, 3, 1}
		want := 0

		got := maxProfit(prices)

		utils.CheckEqual(got, want, t)
	})
}
