package p133

import "terencefan.com/leetcode/src/utils"

type Node = utils.Node

func clone(node *Node, cloned map[int]*Node) *Node {
	if node == nil {
		return nil
	} else if next, ok := cloned[node.Val]; ok {
		return next
	}

	to := &Node{Val: node.Val}
	cloned[to.Val] = to

	if node.Neighbors == nil {
		return to
	}

	to.Neighbors = make([]*Node, len(node.Neighbors))

	for i, neighbor := range node.Neighbors {
		to.Neighbors[i] = clone(neighbor, cloned)
	}

	return to
}

func cloneGraph(node *Node) *Node {
	var cloned = make(map[int]*Node)
	return clone(node, cloned)
}