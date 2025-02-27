package main

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func minInt(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func recursive(node *TreeNode) (min, max, abs int) {
	if node.Left == nil && node.Right == nil {
		return node.Val, node.Val, 0
	} else if node.Left == nil {
		min, max, abs = recursive(node.Right)
	} else if node.Right == nil {
		min, max, abs = recursive(node.Left)
	} else {
		min1, max1, abs1 := recursive(node.Left)
		min2, max2, abs2 := recursive(node.Right)
		min = minInt(min1, min2)
		max = maxInt(max1, max2)
		abs = maxInt(abs1, abs2)
	}
	min = minInt(min, node.Val)
	max = maxInt(max, node.Val)
	abs = maxInt(abs, node.Val - min)
	abs = maxInt(abs, max - node.Val)
	return
}

func maxAncestorDiff(root *TreeNode) int {
	if root == nil {
		return 0
	}
	_, _, c := recursive(root)
	return c
}
