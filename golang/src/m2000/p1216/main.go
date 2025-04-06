package p1216

func isValidPalindrome(s string, k int) bool {
	n := len(s)
	if n < 2 {
		return true
	}

	var dp = make([][]int, n)
	for i := range n {
		dp[i] = make([]int, n+1)
		dp[i][0] = 0
		dp[i][1] = 0
	}

	for k := 2; k <= n; k++ {
		for i := range n - k + 1 {
			if s[i] == s[i+k-1] {
				dp[i][k] = dp[i+1][k-2]
			} else {
				dp[i][k] = min(dp[i][k-1], dp[i+1][k-1]) + 1
			}
		}
	}
	return dp[0][n] <= k
}
