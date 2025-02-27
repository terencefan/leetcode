package main

func tribonacci(n int) int {
	var m = []int{0, 1, 1}

	if n < 3 {
		return m[n]
	}

	for i := 3; i <= n; i++ {
		m[i%3] = m[0] + m[1] + m[2]
	}
	return m[n%3]
}
