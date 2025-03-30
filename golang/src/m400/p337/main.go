package p337

import (
	"terencefan.com/leetcode/src/utils"
)

type TreeNode = utils.TreeNode

func dfs(node *TreeNode) (robbed, untouched int) {
	if node == nil {
		return 0, 0
	}

	leftRobbed, leftUntouched := dfs(node.Left)
	rightRobbed, rightUntouched := dfs(node.Right)

	return leftUntouched + rightUntouched + node.Val, max(leftRobbed, leftUntouched) + max(rightRobbed, rightUntouched)
}

func rob(root *TreeNode) int {
	robbed, untouched := dfs(root)
	return max(robbed, untouched)
}
