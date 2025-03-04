package p97

func isInterleave(s1 string, s2 string, s3 string) bool {
	var m, n = len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	var dp = make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	for i := 1; i <= m; i++ {
		if s1[i-1] == s3[i-1] {
			dp[i][0] = true
		} else {
			break
		}
	}
	for i := 1; i <= n; i++ {
		if s2[i-1] == s3[i-1] {
			dp[0][i] = true
		} else {
			break
		}
	}
	dp[0][0] = true

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if s1[i-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}
			if s2[j-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}
	}
	return dp[m][n]
}
