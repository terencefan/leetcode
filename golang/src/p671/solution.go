package main

func findSecondMinimumValue(root *TreeNode) (r int) {
	if root == nil {
		return -1
	}
	r = -1

	s := []*TreeNode{root}
	for len(s) > 0 {
		node := s[0]
		if node.Val != root.Val && (node.Val < r || r == -1) {
			r = node.Val
		}
		if node.Left != nil {
			s = append(s, node.Left)
		}
		if node.Right != nil {
			s = append(s, node.Right)
		}
		s = s[1:]
	}
	return
}
