package p78

import "slices"

func dfs(r *[][]int, cur, nums []int, start int) {
	if start == len(nums) {
		*r = append(*r, slices.Clone(cur))
		return
	}

	dfs(r, cur, nums, start+1)
	dfs(r, append(cur, nums[start]), nums, start+1)
}

func subsets(nums []int) [][]int {
	var r = make([][]int, 0)
	dfs(&r, []int{}, nums, 0)
	return r
}
