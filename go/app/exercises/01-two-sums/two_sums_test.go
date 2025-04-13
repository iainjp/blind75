package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/utils"
)

var testCases = []struct {
	input  []int
	target int
	want   []int
}{
	{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	{[]int{3, 2, 4}, 6, []int{1, 2}},
	{[]int{3, 3}, 6, []int{0, 1}},
}

func TestTwoSums(t *testing.T) {

	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := twoSum(test.input, test.target)
			utils.CheckEqual(got, test.want, t)
		})
	}
}

func BenchmarkTwoSums(b *testing.B) {
	for b.Loop() {
		_ = twoSum(testCases[0].input, testCases[0].target)
	}
}
