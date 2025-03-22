package p863

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func traverse(parents *map[*TreeNode]*TreeNode, node *TreeNode) {
	if node.Left != nil {
		(*parents)[node.Left] = node
		traverse(parents, node.Left)
	}
	if node.Right != nil {
		(*parents)[node.Right] = node
		traverse(parents, node.Right)
	}
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	var parents = make(map[*TreeNode]*TreeNode)
	var visited = make(map[*TreeNode]bool)

	traverse(&parents, root)
	visited[root] = true

	step := 0
	var q = make([]*TreeNode, 0)

	for len(q) > 0 {
		l := len(q)

		for i := range l {
			node := q[i]
			if step == k {
				var r = make([]int, 0)
				r = append(r, node.Val)
			}
		}

		if len(r) > 0 {
			return r
		}
	}

}
