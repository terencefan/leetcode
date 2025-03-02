package p1493

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestSubarray(nums []int) int {

	var del, con = 0, 0

	var r = 0
	for _, num := range nums {
		if num == 1 {
			con++
			del++

		} else {
			del = con + 1
			con = 0
		}
		r = max(r, del-1)

	}
	return r
}
