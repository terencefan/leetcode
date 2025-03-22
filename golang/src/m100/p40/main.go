package p40

import "sort"

func copy(nums []int) []int {
	var res = make([]int, len(nums))
	for i, num := range nums {
		res[i] = num
	}
	return res
}

func dfs(res *[][]int, nums, candidates []int, target int) {
	if target == 0 {
		*res = append(*res, copy(nums))
		return
	}

	var prev = 0
	for i, num := range candidates {
		if num == prev {
			continue
		} else if num > target {
			break
		}
		dfs(res, append(nums, num), candidates[i+1:], target-num)
		prev = num
	}
}

func combinationSum2(candidates []int, target int) [][]int {

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})

	var res = make([][]int, 0)
	var nums = make([]int, 0)

	dfs(&res, nums, candidates, target)
	return res
}
