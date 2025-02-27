package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

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

const INT_MIN = -(1 << 31)

func MaxSumSubmatrix(matrix [][]int, k int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	var height, width = len(matrix), len(matrix[0])

	var r = -INT_MIN

	for i := 0; i < width; i++ {
		var s = make([]int, height)
		for j := i; j < width; j++ {
			for q := 0; q < height; q++ {
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
