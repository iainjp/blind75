package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/utils"
)

func TestLongestSubstring(t *testing.T) {
	type TestCase struct {
		name  string
		input string
		want  int
	}

	cases := []TestCase{
		{name: "One", input: "abcabcbb", want: 3},
		{name: "Two", input: "bbbbb", want: 1},
		{name: "Three", input: "pwwkew", want: 3},
	}

	for _, test := range cases {
		testName := fmt.Sprintf("Case [%v]", test.name)
		t.Run(testName, func(t *testing.T) {
			got := lengthOfLongestSubstring(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}
