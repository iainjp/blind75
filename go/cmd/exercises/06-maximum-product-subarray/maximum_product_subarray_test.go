package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/cmd/utils"
)

func TestMaximumProductSubarray(t *testing.T) {
	var testCases = []struct {
		input []int
		want  int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
	}

	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := maxProduct(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}
