package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/utils"
)

var testCases = []struct {
	input []int
	want  []int
}{
	{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
	{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
}

func TestProductExceptSelf(t *testing.T) {
	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := productExceptSelf(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}

func BenchmarkProductExceptSelf(b *testing.B) {
	for b.Loop() {
		_ = productExceptSelf(testCases[0].input)
	}
}
