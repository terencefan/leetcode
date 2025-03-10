package p98

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func isValid(node *TreeNode) (bool, int, int) {
	l, r := node.Val, node.Val
	if node.Left != nil {
		lb, ll, lr := isValid(node.Left)
		if !lb || lr >= node.Val {
			return false, -1, -1
		}
		l = ll
	}
	if node.Right != nil {
		rb, rl, rr := isValid(node.Right)
		if !rb || rl <= node.Val {
			return false, -1, -1
		}
		r = rr
	}
	return true, l, r
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	r, _, _ := isValid(root)
	return r
}
