package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/utils"
)

var testCases = []struct {
	input []int
	want  int
}{
	{[]int{7, 1, 5, 3, 6, 4}, 5},
	{[]int{7, 6, 4, 3, 1}, 0},
}

func TestTimeToBuyStocks(t *testing.T) {
	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := maxProfit(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}

func BenchmarkTimeToBuyStocks(b *testing.B) {
	for b.Loop() {
		_ = maxProfit(testCases[0].input)
	}
}
