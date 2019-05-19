package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	t := invertTree(root.Right)
	root.Right = invertTree(root.Left)
	root.Left = t
	return root
}
