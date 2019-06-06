package main

func trailingZeroes(n int) int {
	if n < 5 {
		return 0
	}

	r, i := 0, 5
	for i < n {
		r += n / i
		i *= 5
	}
	return r
}
