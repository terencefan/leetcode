package p637

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func averageOfLevels(root *TreeNode) []float64 {
	q := []*TreeNode{root}

	var r = make([]float64, 0)

	for len(q) > 0 {
		q1 := make([]*TreeNode, 0)

		sum, count := 0, 0
		for _, node := range q {
			if node == nil {
				continue
			}
			sum += node.Val
			count++
			q1 = append(q1, node.Left, node.Right)
		}

		if count > 0 {
			r = append(r, float64(sum)/float64(count))
		}
		q = q1
	}
	return r
}
