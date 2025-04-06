package p1530

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func find(node *TreeNode, distance int) (int, map[int]int) {
	if node == nil {
		return 0, map[int]int{}
	} else if node.Left == nil && node.Right == nil {
		return 0, map[int]int{0: 1}
	}

	c1, m1 := find(node.Left, distance)
	c2, m2 := find(node.Right, distance)

	c := c1 + c2
	m := make(map[int]int)
	for i := range distance - 1 {
		for j := range distance - i - 1 {
			c += m1[i] * m2[j]
		}
		m[i+1] = m1[i] + m2[i]
	}
	return c, m
}

func countPairs(root *TreeNode, distance int) int {
	if root == nil {
		return 0
	}

	c, _ := find(root, distance)
	return c
}
