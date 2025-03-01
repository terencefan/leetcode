package p198

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	var rob1, rob2 = nums[0], 0
	for i := 1; i < len(nums); i++ {
		var current = max(rob1, rob2+nums[i])
		rob2 = rob1
		rob1 = current
	}
	return rob1
}
