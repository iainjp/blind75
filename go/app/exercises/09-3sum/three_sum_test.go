package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/utils"
)

var testCases = []struct {
	input []int
	want  [][]int
}{
	{[]int{-1, 0, 1, 2, -1, -4}, [][]int{{-1, -1, 2}, {-1, 0, 1}}},
	{[]int{0, 1, 1}, [][]int{}},
	{[]int{0, 0, 0}, [][]int{{0, 0, 0}}},
}

func TestSearchRotatedArray(t *testing.T) {
	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := threeSum(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}

func BenchmarkSearchRotatedArray(b *testing.B) {
	for b.Loop() {
		_ = threeSum(testCases[0].input)
	}
}
