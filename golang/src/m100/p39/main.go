package p39

func dfs(r *[][]int, nums *[]int, comb []int, target int) {
	if target == 0 {
		*r = append(*r, append([]int{}, *nums...))
		return
	} else if len(comb) == 0 {
		return
	}

	val := comb[0]
	for i := 0; i*val <= target; i++ {
		dfs(r, nums, comb[1:], target-i*val)
		*nums = append(*nums, val)
	}
	for len(*nums) > 0 && (*nums)[len(*nums)-1] == val {
		*nums = (*nums)[:len(*nums)-1]
	}
}

func combinationSum(candidates []int, target int) [][]int {
	var r = make([][]int, 0)
	var nums = make([]int, 0)
	dfs(&r, &nums, candidates, target)
	return r
}
