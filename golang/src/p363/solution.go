package main

func Kadane(a []int, k int) int {
	var n = len(a)
	var r = -99999

	for i := 0; i < n; i++ {
		var s = 0
		for j := i; j < n; j++ {
			s += a[j]
			if s <= k && s > r {
				r = s
			}
		}
	}
	return r
}

func MaxSumSubmatrix(matrix [][]int, k int) int {
	var m = len(matrix)
	var n = len(matrix[0])
	var r = -99999

	for i := 0; i < n; i++ {
		var s = make([]int, m)
		for p := 0; p < m; p++ {
			s[p] = 0
		}
		for j := i; j < n; j++ {
			for q := 0; q < m; q++ {
				s[q] += matrix[q][j]
			}
			var x = Kadane(s, k)
			if x > r && x <= k {
				r = x
			}
		}
	}
	return r
}
