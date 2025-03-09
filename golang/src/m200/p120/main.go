package p120

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	n := len(triangle)

	dp := make([]int, n)
	dp[0] = triangle[0][0]

	for i := range triangle {
		if i == 0 {
			continue
		}

		dp1 := make([]int, n)
		for k := 0; k <= i; k++ {
			if k == 0 {
				dp1[k] = dp[k] + triangle[i][k]
			} else if k == i {
				dp1[k] = dp[k-1] + triangle[i][k]
			} else {
				dp1[k] = min(dp[k-1], dp[k]) + triangle[i][k]
			}
		}
		dp = dp1
	}

	r := 1 << 31
	for i := range dp {
		r = min(r, dp[i])
	}
	return r
}
