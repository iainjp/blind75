package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/utils"
)

var testCases = []struct {
	input [][]int
	want  [][]int
}{
	{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}},
	{[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}},
}

func TestMergeInverval(t *testing.T) {
	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := merge(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}

func BenchmarkMergeInterval(b *testing.B) {
	for b.Loop() {
		_ = merge(testCases[0].input)
	}
}
