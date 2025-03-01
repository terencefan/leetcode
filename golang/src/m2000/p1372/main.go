package p1372

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

const (
	Left = iota
	Right
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longest(node *TreeNode, direction int) (int, int) {
	if node == nil {
		return 0, 0
	}

	if direction == Left {
		next, maxRight := longest(node.Right, Right)
		_, maxLeft := longest(node.Left, Left)
		return next + 1, max(next+1, max(maxLeft, maxRight))
	} else {
		next, maxLeft := longest(node.Left, Left)
		_, maxRight := longest(node.Right, Right)
		return next + 1, max(next+1, max(maxLeft, maxRight))
	}
}

func longestZigZag(root *TreeNode) int {
	var r = 0
	if root.Left != nil {
		_, maxLeft := longest(root.Left, Left)
		r = max(r, maxLeft)
	}
	if root.Right != nil {
		_, maxRight := longest(root.Right, Right)
		r = max(r, maxRight)
	}
	return r
}
