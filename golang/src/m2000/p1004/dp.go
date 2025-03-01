package p1004

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func longestOnes(nums []int, k int) int {
	var dp = make([]int, k+1)

	var r = 0
	for _, num := range nums {
		if num == 1 {
			for j := range k + 1 {
				dp[j]++
			}
		} else {
			for j := k; j > 0; j-- {
				dp[j] = dp[j-1] + 1
			}
			dp[0] = 0
		}
		r = max(r, dp[k])
	}
	return r
}
