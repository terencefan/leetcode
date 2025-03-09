package p103

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func reverse(s []int) []int {
	l, r := 0, len(s)-1
	for l < r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return s
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	leftToRight := true

	q := []*TreeNode{root}
	r := make([][]int, 0)
	for len(q) > 0 {
		r1 := make([]int, 0)

		l := len(q)
		for i := range l {
			node := q[i]
			if node == nil {
				continue
			}
			q = append(q, node.Left, node.Right)
			r1 = append(r1, node.Val)
		}
		q = q[l:]

		if len(r1) > 0 {
			if leftToRight {
				r = append(r, r1)
			} else {
				r = append(r, reverse(r1))
			}
		}
		leftToRight = !leftToRight
	}
	return r
}
