package p45

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func jump(nums []int) int {
	if len(nums) < 2 {
		return 0
	}

	var r = 0
	var cur, farthest = 0, 0

	for i, num := range nums[:len(nums)-1] {
		if farthest < i+num {
			farthest = i + num
		}

		if i == cur {
			r++
			cur = farthest
		}
	}
	return r
}
