package main

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}
	l := len(nums)
	left, right := make([]int, len(nums)), make([]int, len(nums))

	left[0] = nums[0]
	for i := 1; i < l; i++ {
		left[i] = min(left[i - 1], nums[i])
	}

	right[l - 1] = nums[l - 1]
	for i := l - 2; i >= 0; i-- {
		right[i] = max(right[i + 1], nums[i])
	}

	for i := 1; i < l - 1; i++ {
		if left[i - 1] < nums[i] && right[i + 1] > nums[i] {
			return true
		}
	}
	return false
}
