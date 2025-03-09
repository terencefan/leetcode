package p55

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func canJump(nums []int) bool {
	if len(nums) == 0 {
		return true
	}

	cmax := 0
	for i, num := range nums {
		if cmax < i {
			return false
		} else if cmax >= len(nums)-1 {
			return true
		}
		cmax = max(cmax, i+num)
	}
	return false
}
