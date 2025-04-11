package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/cmd/utils"
)

func TestLongestSubstring(t *testing.T) {
	type TestCase struct {
		name  string
		input []int
		want  int
	}

	cases := []TestCase{
		{name: "One", input: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, want: 6},
		{name: "Two", input: []int{1}, want: 1},
		{name: "Three", input: []int{5, 4, -1, 7, 8}, want: 23},
	}

	for _, test := range cases {
		testName := fmt.Sprintf("Case [%v]", test.name)
		t.Run(testName, func(t *testing.T) {
			got := maxSubArray(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}
