package exercises

import (
	"fmt"
	"testing"

	"github.com/iainjp/blind75/utils"
)

var testCases = []struct {
	input [][]int
}{
	{[][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}}},
	{[][]int{{}}},
	{[][]int{}},
}

func deserialise(s [][]int) *Node {
	if len(s) == 0 {
		return nil
	}

	nodes := make(map[int]*Node)

	// populate nodes map first
	for i := 1; i <= len(s); i++ {
		n := Node{
			Val:       i, // 1-indexed
			Neighbors: []*Node{},
		}
		nodes[i] = &n
	}

	// populate node neighbours
	for val, neighbours := range s {
		i := val + 1
		for _, n := range neighbours {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[n])
		}
	}

	return nodes[1]
}

// func serialise(graph *Node) [][]int {
// 	return [][]int{}
// }

func TestCloneGraph(t *testing.T) {
	for testNum, test := range testCases {
		testName := fmt.Sprintf("Case [%v]", testNum)
		t.Run(testName, func(t *testing.T) {
			got := cloneGraph(deserialise(test.input))
			utils.CheckEqual(got, deserialise(test.input), t)
		})
	}
}

func BenchmarkCloneGraph(b *testing.B) {
	for b.Loop() {
		_ = cloneGraph(deserialise(testCases[0].input))
	}
}
