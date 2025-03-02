package p746

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCostClimbingStairs(cost []int) int {

	if len(cost) == 0 {
		return 0
	} else if len(cost) == 1 {
		return cost[0]
	}

	dp1, dp2 := 0, 0

	for i := 2; i <= len(cost); i++ {
		dp1, dp2 = min(dp1+cost[i-1], dp2+cost[i-2]), dp1
	}

	return dp1
}
