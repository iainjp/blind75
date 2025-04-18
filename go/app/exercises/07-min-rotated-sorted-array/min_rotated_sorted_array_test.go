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
	{[]int{3, 4, 5, 1, 2}, 1},
	{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
	{[]int{11, 13, 15, 17}, 11},
}

func TestMinRotatedSortedArray(t *testing.T) {
	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := findMin(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}

func BenchmarkMinRotatedSortedArray(b *testing.B) {
	for b.Loop() {
		_ = findMin(testCases[0].input)
	}
}
