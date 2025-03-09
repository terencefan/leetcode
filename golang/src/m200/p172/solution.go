package p172

func trailingZeroes(n int) int {
	x := 5

	var r = 0
	for n >= x {
		r += n / x
		x *= 5
	}
	return r
}
