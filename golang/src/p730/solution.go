package main

func countPalindromicSubsequences(S string) int {
	dp := make([][][]int, 26)

	for i := 0; i < 26; i++ {
		dp[i] = make([][]int, len(S))
		for j := 0; j < len(S); j++ {
		}
	}
	return 1
}


