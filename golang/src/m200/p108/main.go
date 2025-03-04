package p108

import (
	"terencefan.com/leetcode/src/utils"
)

type TreeNode = utils.TreeNode

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	} else if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}

	mid := len(nums) / 2

	return &TreeNode{
		Val:   nums[mid],
		Left:  sortedArrayToBST(nums[:mid]),
		Right: sortedArrayToBST(nums[mid+1:]),
	}
}
