package p216

func dfs(k, n, s int, nums []int, res *[][]int) {
	if k == 0 {
		if n == 0 {
			dst := make([]int, len(nums))
			copy(dst, nums)
			*res = append(*res, dst)
			return
		}
		return
	}

	for num := s; num < 10; num++ {
		if num > n {
			break
		}
		nums[k-1] = num
		dfs(k-1, n-num, num+1, nums, res)
	}
}

func combinationSum3(k int, n int) [][]int {
	var res = make([][]int, 0)
	var nums = make([]int, k)

	dfs(k, n, 1, nums, &res)
	return res
}
