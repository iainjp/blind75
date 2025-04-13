package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/utils"
)

func TestMaximumSubarray(t *testing.T) {
	var testCases = []struct {
		input []int
		want  int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}

	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := maxSubArray(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}
