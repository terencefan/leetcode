package p189

func Rotate(nums []int, k int) {
	var l = len(nums)
	k %= len(nums)
	if k == 0 || l == 1 {
		return
	}

	var t = make([]int, k)
	copy(t, nums[l-k:])
	copy(nums[k:], nums[:l-k])
	copy(nums, t)
}
