package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/utils"
)

var testCases = []struct {
	input []int
	want  bool
}{
	{[]int{1, 2, 3, 1}, true},
	{[]int{1, 2, 3, 4}, false},
	{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
}

func TestContainsDuplicates(t *testing.T) {
	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := containsDuplicate(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}

func BenchmarkContainsDuplicates(b *testing.B) {
	for b.Loop() {
		_ = containsDuplicate(testCases[0].input)
	}
}
