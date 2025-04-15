package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/utils"
)

var testCases = []struct {
	input int
	want  int
}{
	{11, 3},
	{128, 1},
	{2147483645, 30},
}

func TestHammingWeight(t *testing.T) {
	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := hammingWeight(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}

func BenchmarkHammingWeight(b *testing.B) {
	for b.Loop() {
		_ = hammingWeight(testCases[0].input)
	}
}
