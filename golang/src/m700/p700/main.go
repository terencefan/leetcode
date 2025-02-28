package main

import "terencefan.com/leetcode/src/utils"

type TreeNode utils.TreeNode

func searchBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}

	var node = searchBST((*TreeNode)(root.Left), val)
	if node != nil {
		return node
	}
	return searchBST((*TreeNode)(root.Right), val)
}
