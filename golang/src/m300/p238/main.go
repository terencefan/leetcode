package p238

func productExceptSelf(nums []int) []int {
	var r = make([]int, len(nums))

	r[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		r[i] = r[i-1] * nums[i]
	}

	r[len(nums)-1] = r[len(nums)-2]
	for i := len(nums) - 2; i >= 1; i-- {
		r[i] = r[i-1] * nums[i+1]
		nums[i] *= nums[i+1]
	}
	r[0] = nums[1]
	return r
}
