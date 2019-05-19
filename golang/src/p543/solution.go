package p543

import (
	"fmt"
)

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func recursive(node *TreeNode) (r, l int) {
	if node == nil {
		return 0, 0
	} else {
		r1, l1 := recursive(node.Left)
		r2, l2 := recursive(node.Right)
		return max(max(r1, r2), l1 + l2 + 1), max(l1, l2) + 1
	}
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r, _ := recursive(root)
	return r - 1
}

func main() {
	n5 := &TreeNode{1, nil, nil}
	n4 := &TreeNode{1, nil, nil}
	n3 := &TreeNode{1, nil, nil}
	n2 := &TreeNode{1, n4, n5}
	n1 := &TreeNode{1, n2, n3}
	r := diameterOfBinaryTree(n1)
	fmt.Println(r)
}