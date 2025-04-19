package exercises

type Node struct {
	Val       int
	Neighbors []*Node
}

func cloneGraph(node *Node) *Node {
	// short-circuit nil case
	if node == nil {
		return nil
	}

	// short-circuit single node case
	if len(node.Neighbors) == 0 {
		return &Node{
			Val:       node.Val,
			Neighbors: []*Node{},
		}
	}

	// walk the graph, cloning each node, add to map
	// then walk again and assign neighbours
	// -- technically O(n) (cancel the 2*)
	// blowing up the memory use with multiple maps though

	clonedNodes := make(map[int]*Node)
	haveSeen := make(map[int]bool)
	haveSetNeighbours := make(map[int]bool)

	clone := func(n *Node) {
		clonedNodes[n.Val] = &Node{
			Val: n.Val,
		}
	}

	setNeighbours := func(n *Node) {
		cloned := clonedNodes[n.Val]
		for _, neighbour := range n.Neighbors {
			cloned.Neighbors = append(cloned.Neighbors, clonedNodes[neighbour.Val])
		}

		clonedNodes[n.Val] = cloned
	}

	var walk func(n *Node, cb func(n *Node), tracker map[int]bool)
	walk = func(n *Node, cb func(n *Node), tracker map[int]bool) {
		_, seen := tracker[n.Val]

		// haven't seen it yet
		if !seen {
			tracker[n.Val] = true
			cb(n)
			for _, n := range n.Neighbors {
				walk(n, cb, tracker)
			}
		}
	}

	walk(node, clone, haveSeen)
	walk(node, setNeighbours, haveSetNeighbours)

	return clonedNodes[1]
}
