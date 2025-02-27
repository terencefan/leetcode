package main

import "leetcode"

type TreeNode leetcode.TreeNode

func recursive(node *TreeNode, level int, r *[]int) {
	if node == nil {
		return
	}
	if level >= len(*r) {
		*r = append(*r, node.Val)
	}
	recursive((*TreeNode)(node.Right), level+1, r)
	recursive((*TreeNode)(node.Left), level+1, r)
}

func rightSideViewRecursive(root *TreeNode) []int {
	var r = make([]int, 0)
	recursive(root, 0, &r)
	return r
}

type NodeLevel struct {
	node  *TreeNode
	level int
}

func rightSideView(root *TreeNode) []int {
	var r = make([]int, 0)
	var stack = make([]NodeLevel, 0)

	stack = append(stack, NodeLevel{root, 0})

	for len(stack) > 0 {
		item := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node, level := item.node, item.level

		if node == nil {
			continue
		}
		if level >= len(r) {
			r = append(r, node.Val)
		}
		stack = append(stack, NodeLevel{(*TreeNode)(node.Left), level + 1})
		stack = append(stack, NodeLevel{(*TreeNode)(node.Right), level + 1})
	}
	return r
}
