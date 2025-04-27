package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/utils"
)

var testCases = []struct {
	existingInput [][]int
	newInput      []int
	want          [][]int
}{
	{[][]int{{1, 3}, {6, 9}}, []int{2, 5}, [][]int{{1, 5}, {6, 9}}},
	{[][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}}, []int{4, 8}, [][]int{{1, 2}, {3, 10}, {12, 16}}},
}

func TestInsertInverval(t *testing.T) {
	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := insert(test.existingInput, test.newInput)
			utils.CheckEqual(got, test.want, t)
		})
	}
}

func BenchmarkInsertInterval(b *testing.B) {
	for b.Loop() {
		_ = insert(testCases[0].existingInput, testCases[0].newInput)
	}
}
