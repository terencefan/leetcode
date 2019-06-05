package main

func min(a, b int) int {
	if a == 0 {
		return b
	} else if a < b {
		return a
	} else {
		return b
	}
}

func jump_dp(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 0

	for i, num := range nums {
		for j := 1; j <= num && i + j < len(nums); j++ {
			dp[i + j] = min(dp[i + j], dp[i] + 1)
		}
	}
	return dp[len(nums) - 1]
}
