package p11

func maxArea(heights []int) int {
	area := 0
	l, r := 0, len(heights)-1
	for l < r {
		area = max(area, min(heights[l], heights[r])*(r-l))
		if heights[l] < heights[r] {
			l++
		} else {
			r--
		}
	}
	return area
}
