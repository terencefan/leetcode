package p338

func countBits(n int) []int {
	var r = make([]int, n+1)
	if n == 0 {
		return []int{0}
	} else if n == 1 {
		return []int{1}
	}

	r[0] = 0
	r[1] = 1

	x := 2

	for i := 2; i <= n; i++ {
		r[i] = r[i-x] + 1
		if i == x*2-1 {
			x *= 2
		}
	}
	return r
}
