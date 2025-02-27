package main

func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	s := []*TreeNode{}

	node := root
	for node.Left != nil {
		s = append(s, node)
		node = node.Left
	}
	s = append(s, node)
	root = node

	for i := len(s) - 1; i > 0; i-- {
		parent, current := s[i-1], s[i]
		current.Right = parent
		current.Left = parent.Right
	}
	return root
}
