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
	{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	{[]int{1, 1}, 1},
}

func TestContainerWithMostWater(t *testing.T) {
	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := maxArea(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}

func BenchmarkContainerWithMostWater(b *testing.B) {
	for b.Loop() {
		_ = maxArea(testCases[0].input)
	}
}
