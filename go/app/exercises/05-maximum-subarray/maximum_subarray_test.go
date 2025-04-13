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
	{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	{[]int{1}, 1},
	{[]int{5, 4, -1, 7, 8}, 23},
}

func TestMaximumSubarray(t *testing.T) {
	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := maxSubArray(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}

func BenchmarkMaximumSubarray(b *testing.B) {
	for b.Loop() {
		_ = maxSubArray(testCases[0].input)
	}
}
