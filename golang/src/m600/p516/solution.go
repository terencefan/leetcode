package main

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestPalindromeSubseq(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	dp := make([][]int, len(s))
	for i := range dp {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
	}

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			dp[i][i+1] = 2
		} else {
			dp[i][i+1] = 1
		}
	}

	for k := 2; k < len(s); k++ {
		for i := 0; i < len(s)-k; i++ {
			dp[i][i+k] = max(dp[i+1][i+k], dp[i][i+k-1])
			if s[i] == s[i+k] {
				dp[i][i+k] = max(dp[i][i+k], dp[i+1][i+k-1]+2)
			}
		}
	}
	return dp[0][len(s)-1]
}
