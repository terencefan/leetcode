package p112

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == sum
	}
	if root.Left != nil && hasPathSum(root.Left, sum-root.Val) {
		return true
	}
	if root.Right != nil && hasPathSum(root.Right, sum-root.Val) {
		return true
	}
	return false
}
