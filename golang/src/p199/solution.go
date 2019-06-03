package main

func rightSideView(root *TreeNode) (r []int) {
	if root == nil {
		return []int{}
	}

	q := []*TreeNode{root}

	for len(q) > 0 {
		l := len(q)
		r = append(r, q[len(q) - 1].Val)
		for i := 0; i < l; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[l:]
	}
	return r
}
