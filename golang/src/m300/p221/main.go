package p224

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
	}

	for i := range m {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
		}
	}
	for i := range n {
		if matrix[0][i] == '1' {
			dp[0][i] = 1
		}
	}

	for i := range m {
		for j := range n {
			if i < 1 || j < 1 {
				continue
			}
			if matrix[i][j] == '0' {
				dp[i][j] = 0
			} else {
				left, up := dp[i-1][j], dp[i][j-1]
				if left == up {
					if matrix[i-left][j-left] == '1' {
						dp[i][j] = left + 1
					} else {
						dp[i][j] = left
					}
				} else {
					dp[i][j] = min(left, up) + 1
				}
			}
		}
	}

	var r = 0
	for i := range m {
		for j := range n {
			if dp[i][j] > r {
				r = dp[i][j]
			}
		}
	}
	return r * r
}
