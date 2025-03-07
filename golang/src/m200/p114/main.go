package p114

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	s := []*TreeNode{root}

	var prev *TreeNode
	for len(s) > 0 {
		node := s[len(s)-1]
		s = s[:len(s)-1]

		if prev != nil {
			prev.Right = node
			prev.Left = nil
		}
		prev = node

		if node.Right != nil {
			s = append(s, node.Right)
		}
		if node.Left != nil {
			s = append(s, node.Left)
		}
	}
	prev.Left = nil
}
