package p962

func maxWidthRamp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)

	rightMax := make([]int, len(nums))

	rightMax[n-1] = nums[n-1]

	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], nums[i])
	}

	l, r, s := 0, 0, 0

	for r < n {
		for nums[l] > rightMax[r] {
			l++
		}
		s = max(s, r-l)
		r++
	}
	return s
}
