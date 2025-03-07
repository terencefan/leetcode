package p202

func convert(n int) int {
	r := 0
	for n > 0 {
		r += (n % 10) * (n % 10)
		n /= 10
	}
	return r
}

func isHappy(n int) bool {
	m := make(map[int]bool)
	m[n] = true

	for {
		n = convert(n)
		if n == 1 {
			return true
		} else if m[n] {
			return false
		}
		m[n] = true
	}
}
