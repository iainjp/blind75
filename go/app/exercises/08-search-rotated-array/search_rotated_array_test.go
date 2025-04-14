package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/utils"
)

var testCases = []struct {
	input  []int
	target int
	want   int
}{
	{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
	{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
	{[]int{1}, 0, -1},
}

func TestSearchRotatedArray(t *testing.T) {
	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := search(test.input, test.target)
			utils.CheckEqual(got, test.want, t)
		})
	}
}

func BenchmarkSearchRotatedArray(b *testing.B) {
	for b.Loop() {
		_ = search(testCases[0].input, testCases[0].target)
	}
}
