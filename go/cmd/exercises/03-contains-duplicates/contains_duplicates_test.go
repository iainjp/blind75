package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/cmd/utils"
)

func TestContainsDuplicates(t *testing.T) {
	var testCases = []struct {
		input []int
		want  bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := containsDuplicate(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}
