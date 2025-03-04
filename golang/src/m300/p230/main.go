package p230

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func findKth(node *TreeNode, k int) (int, bool) {
	if node == nil || k <= 0 {
		return 0, false
	}

	v1, ok := findKth(node.Left, k)
	if ok {
		return v1, true
	} else if v1 == k-1 {
		return node.Val, true
	}

	v2, ok := findKth(node.Right, k-v1-1)
	if ok {
		return v2, true
	}
	return v1 + 1 + v2, false
}

func kthSmallest(root *TreeNode, k int) int {
	if val, ok := findKth(root, k); ok {
		return val
	}
	return -1
}
