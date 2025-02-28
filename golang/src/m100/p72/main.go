package p72

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func min3(a, b, c int) int {
	return min(min(a, b), c)
}

func minDistance(word1 string, word2 string) int {
	len1, len2 := len(word1), len(word2)

	var dp = make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}

	for i := range len1 {
		dp[i+1][0] = i + 1
	}
	for j := range len2 {
		dp[0][j+1] = j + 1
	}
	dp[0][0] = 0

	for i := range len1 {
		for j := range len2 {
			if word1[i] == word2[j] {
				dp[i+1][j+1] = dp[i][j]
			} else {
				dp[i+1][j+1] = min3(dp[i][j], dp[i+1][j], dp[i][j+1]) + 1
			}
		}
	}
	return dp[len1][len2]
}
