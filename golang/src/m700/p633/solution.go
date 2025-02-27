package main

func judgeSquareSum(c int) bool {
	m := make(map[int]bool)
	for i := 0; i*i <= c; i++ {
		m[i*i] = true
		if m[c-i*i] {
			return true
		}
	}
	return false
}
