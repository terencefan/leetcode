package main

import "terencefan.com/leetcode/src/utils"

type TreeNode utils.TreeNode

func dfs(node *TreeNode, level int, levelSum *[]int) {
	if node == nil {
		return
	}

	for len(*levelSum) < level+1 {
		*levelSum = append(*levelSum, 0)
	}
	(*levelSum)[level] += node.Val
	dfs((*TreeNode)(node.Left), level+1, levelSum)
	dfs((*TreeNode)(node.Right), level+1, levelSum)
}

func maxLevelSum(root *TreeNode) int {
	var levelSum = make([]int, 0)

	dfs(root, 0, &levelSum)

	r, maxv := 0, -1<<31
	for i, v := range levelSum {
		if v > maxv {
			maxv = v
			r = i
		}
	}
	return r + 1
}
