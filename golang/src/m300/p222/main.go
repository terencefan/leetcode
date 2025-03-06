package p222

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	q := []*TreeNode{root}

	count := 0
	for len(q) > 0 {
		l := len(q)
		for i := 0; i < l; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[l:]
		count += l
	}
	return count
}
