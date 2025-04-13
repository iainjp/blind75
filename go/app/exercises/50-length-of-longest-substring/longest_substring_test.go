package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/utils"
)

var testCases = []struct {
	input string
	want  int
}{
	{"abcabcbb", 3},
	{"bbbbb", 1},
	{"pwwkew", 3},
}

func TestLongestSubstring(t *testing.T) {
	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := lengthOfLongestSubstring(test.input)
			utils.CheckEqual(got, test.want, t)
		})
	}
}

func BenchmarkLongestSubstring(b *testing.B) {
	for b.Loop() {
		_ = lengthOfLongestSubstring(testCases[0].input)
	}
}
