package p102

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func levelOrder(root *TreeNode) [][]int {
	q := []*TreeNode{root}

	var r = make([][]int, 0)

	for len(q) > 0 {
		q1 := make([]*TreeNode, 0)
		r1 := make([]int, 0)
		for _, node := range q {
			if node == nil {
				continue
			}
			q1 = append(q1, node.Left, node.Right)
			r1 = append(r1, node.Val)
		}
		q = q1
		if len(r1) > 0 {

			r = append(r, r1)
		}
	}
	return r
}
