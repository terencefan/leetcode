package main

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var g = 0

func recursive(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	recursive(root.Right)
	g += root.Val
	root.Val = g
	recursive(root.Left)
	return root
}

func convertBST(root *TreeNode) *TreeNode {
	g = 0
	return recursive(root)
}

func main() {
	n3 := &TreeNode{3, nil, nil}
	n2 := &TreeNode{1, nil, nil}
	n1 := &TreeNode{2, n2, n3}
	// n5 := &TreeNode{1, nil, nil}
	// n4 := &TreeNode{-4, nil, nil}
	// n3 := &TreeNode{3, nil, nil}
	// n2 := &TreeNode{0, n4, n5}
	// n1 := &TreeNode{2, n2, n3}
	r := convertBST(n1)
	fmt.Println(r)
}
