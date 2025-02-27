package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func recursive(root *TreeNode) (a, b int) {
	if root == nil {
		return 0, 0
	}

	a1, b1 := recursive(root.Left)
	a2, b2 := recursive(root.Right)

	b = a1 + a2
	a = max(root.Val+b1+b2, b)
	return
}

func rob(root *TreeNode) int {
	var a, b = recursive(root)
	return max(a, b)
}
