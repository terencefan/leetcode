package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	var s = make([]*TreeNode, 0)

	node := root
	for node != p {
		if p.Val > node.Val {
			node = node.Right
		} else {
			s = append(s, node)
			node = node.Left
		}
	}

	if node.Right != nil {
		node = node.Right
		for node != nil {
			s = append(s, node)
			node = node.Left
		}
	}

	if len(s) > 0 {
		return s[len(s) - 1]
	} else {
		return nil
	}
}