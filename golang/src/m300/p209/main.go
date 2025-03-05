package p209

func getNum(nums []int, i int) int {
	if i < 0 {
		return 0
	}
	return nums[i]
}

func minSubArrayLen(target int, nums []int) int {
	for i := range nums {
		if i > 0 {
			nums[i] += nums[i-1]
		}
	}

	i := -1
	for j := range nums {
		if nums[j] < target {
			continue
		}
		for nums[j]-getNum(nums, i) > target {
			i++
		}
	}
	return 0
}
