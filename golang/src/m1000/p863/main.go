package p863

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func traverse(m *map[*TreeNode]*TreeNode, node *TreeNode) {
	if node.Left != nil {
		(*m)[node.Left] = node
		traverse(m, node.Left)
	}
	if node.Right != nil {
		(*m)[node.Right] = node
		traverse(m, node.Right)
	}
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var parents = make(map[*TreeNode]*TreeNode)
	var visited = make(map[*TreeNode]bool)
	visited[target] = true

	traverse(&parents, root)

	step := 0
	var q = []*TreeNode{target}
	var r = make([]int, 0)

	var appendNode = func(node *TreeNode) {
		if node != nil && !visited[node] {
			visited[node] = true
			q = append(q, node)
		}
	}

	for len(q) > 0 {
		l := len(q)
		for i := range l {
			node := q[i]
			if node == nil {
				continue
			}

			if step == k {
				r = append(r, node.Val)
			} else {
				appendNode(node.Left)
				appendNode(node.Right)
				appendNode(parents[node])
			}
		}
		q = q[l:]
		step++
	}
	return r
}
