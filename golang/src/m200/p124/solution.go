package main

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

var r int

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxChainSum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left, right := maxChainSum(root.Left), maxChainSum(root.Right)
	left = max(0, left)
	right = max(0, right)

	r = max(r, left+right+root.Val)
	return max(left, right) + root.Val
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r = -(1 << 32)
	maxChainSum(root)
	return r
}
