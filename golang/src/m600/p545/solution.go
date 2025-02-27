package main

func boundaryOfBinaryTree(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	var node, leftMost, rightMost *TreeNode

	r := []int{root.Val}

	// find the left boundary
	leftMost, node = root, root.Left
	for node != nil {
		leftMost = node
		r = append(r, node.Val)
		if node.Left != nil {
			node = node.Left
		} else {
			node = node.Right
		}
	}

	// find all leafs
	q := []*TreeNode{root}
	for len(q) > 0 {
		node = q[len(q)-1]
		q = q[:len(q)-1]

		if node.Right == nil && node.Left == nil {
			if node != leftMost {
				rightMost = node
				r = append(r, node.Val)
			}
		}

		if node.Right != nil {
			q = append(q, node.Right)
		}

		if node.Left != nil {
			q = append(q, node.Left)
		}
	}

	s := make([]*TreeNode, 0)

	// find the right boundary
	node = root.Right
	for node != nil {
		if node != rightMost {
			s = append(s, node)
		}
		if node.Right != nil {
			node = node.Right
		} else {
			node = node.Left
		}
	}

	for i := range s {
		k := len(s) - i - 1
		r = append(r, s[k].Val)
	}
	return r
}
