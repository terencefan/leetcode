package p46

func dfs(r *[][]int, nums []int, index int) {
	if index == len(nums)-1 {
		*r = append(*r, append([]int{}, nums...))
	}

	for i := index; i < len(nums); i++ {
		nums[index], nums[i] = nums[i], nums[index]
		dfs(r, nums, index+1)
		nums[index], nums[i] = nums[i], nums[index]
	}
}

func permute(nums []int) [][]int {
	var r = make([][]int, 0)
	dfs(&r, nums, 0)
	return r
}
