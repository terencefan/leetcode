package main

func isLeaf(node *TreeNode) bool {
	return node.Left == nil && node.Right == nil
}

func findLeavesRecursively(node *TreeNode, r *[]int) {
	if node.Left != nil {
		if isLeaf(node.Left) {
			*r = append(*r, node.Left.Val)
			node.Left = nil
		} else {
			findLeavesRecursively(node.Left, r)
		}
	}
	if node.Right != nil {
		if isLeaf(node.Right) {
			*r = append(*r, node.Right.Val)
			node.Right = nil
		} else {
			findLeavesRecursively(node.Right, r)
		}
	}
}

func findLeaves(root *TreeNode) (r [][]int) {
	if root == nil {
		return
	}

	for !isLeaf(root) {
		res := make([]int, 0)
		findLeavesRecursively(root, &res)
		r = append(r, res)
	}
	r = append(r, []int{root.Val})
	return
}
