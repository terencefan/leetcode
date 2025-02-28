package p62

func uniquePaths(m int, n int) int {
	if m == 0 || n == 0 {
		return 0
	}

	if m > n {
		m, n = n, m
	}

	var arr = make([]int, m)
	for i := range arr {
		arr[i] = 1
	}

	for j := 1; j < n; j++ {
		arr[0] = 1
		for i := 1; i < m; i++ {
			arr[i] = arr[i] + arr[i-1]
		}
	}
	return arr[m-1]
}
