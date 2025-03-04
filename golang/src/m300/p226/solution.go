package p226

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}
