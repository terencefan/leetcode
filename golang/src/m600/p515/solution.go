package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func largestValues(root *TreeNode) (r []int) {
	if root == nil {
		return
	}

	q := []*TreeNode{root}
	for len(q) > 0 {
		length, m := len(q), q[0].Val
		for i := 0; i < length; i++ {
			m = max(m, q[i].Val)
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		r = append(r, m)
		q = q[length:]
	}
	return
}
