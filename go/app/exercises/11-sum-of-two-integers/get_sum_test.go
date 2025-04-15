package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/utils"
)

var testCases = []struct {
	a    int
	b    int
	want int
}{
	{1, 2, 3},
	{2, 3, 5},
	{4, 5, 9},
}

func TestContainerWithMostWater(t *testing.T) {
	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := getSum(test.a, test.b)
			utils.CheckEqual(got, test.want, t)
		})
	}
}

func BenchmarkContainerWithMostWater(b *testing.B) {
	for b.Loop() {
		_ = getSum(testCases[0].a, testCases[0].b)
	}
}
