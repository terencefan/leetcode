package codes

func preOrderTraverse(t *TreeNode) (r []int) {
	s := make([]*TreeNode, 0)

	node := t
	for node != nil {
		s = append(s, node)
		node = node.Left
	}

	for len(s) > 0 {
		node := s[len(s)-1]
		s = s[:len(s)-1]

		r = append(r, node.Val)

		if node.Right != nil {
			node = node.Right
			for node != nil {
				r = append(r, node.Val)
				node = node.Left
			}
		}
	}
	return r
}
