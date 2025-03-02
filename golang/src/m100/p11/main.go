package p11

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxArea(height []int) int {
	l, r := 0, len(height)-1

	res := 0
	for l < r {
		product := min(height[l], height[r]) * (r - l)
		if product > res {
			res = product
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}
	return res
}
