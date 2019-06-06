package main

func isCompleteTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var p, q []*TreeNode
	q = append(q, root)

	expected := 1
	for len(q) > 0 {
		length := len(q)
		p, q = q[:length], q[length:]

		for _, node := range p {
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}

		expected <<= 1
		if len(q) != expected {
			consecutive := true
			for _, node := range p {
				if node.Left == nil {
					consecutive = false
				} else {
					if !consecutive {
						return false
					}
				}

				if node.Right == nil {
					consecutive = false
				} else {
					if !consecutive {
						return false
					}
				}
			}

			for _, node := range q {
				if node.Left != nil {
					return false
				}
				if node.Right != nil {
					return false
				}
			}

		}
	}
	return true
}
