package p437

import "terencefan.com/leetcode/src/utils"

type TreeNode = utils.TreeNode

func dfs(node *TreeNode, nums []int, targetSum int) int {
	if node == nil {
		return 0
	}

	var last = 0

	if len(nums) == 0 {
		last = node.Val
	} else {
		last = nums[len(nums)-1] + node.Val
	}
	nums = append(nums, last)

	defer func() {
		nums = nums[:len(nums)-1]
	}()

	var r = 0
	if last == targetSum {
		r++
	}
	for i := range len(nums) - 1 {
		if last-nums[i] == targetSum {
			r++
		}
	}

	return r + dfs(node.Left, nums, targetSum) + dfs(node.Right, nums, targetSum)
}

func pathSum(root *TreeNode, targetSum int) int {
	var nums = make([]int, 0)
	return dfs(root, nums, targetSum)
}
