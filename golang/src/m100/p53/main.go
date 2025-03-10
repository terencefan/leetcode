package p53

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	var r, cmax = nums[0], nums[0]

	for i, num := range nums {
		if i == 0 {
			cmax = num
		} else {
			cmax = max(cmax+num, num)
		}
		r = max(r, cmax)
	}
	return r
}
