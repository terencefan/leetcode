package p64

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

const INTMIN = -1 << 31

func max(nums ...int) int {
	var r = INTMIN
	for _, num := range nums {
		if num > r {
			r = num
		}
	}
	return r
}

func maxPath(node *TreeNode) (include int, exclude int) {
	if node == nil {
		return INTMIN, INTMIN
	}

	li, le := maxPath(node.Left)
	ri, re := maxPath(node.Right)

	include = node.Val + max(0, li, ri)
	exclude = max(node.Val+li+ri, li, le, ri, re)
	return
}

func maxPathSum(root *TreeNode) int {
	x, y := maxPath(root)
	return max(x, y)
}
