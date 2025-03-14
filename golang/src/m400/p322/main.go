package p322

import "sort"

const INTMAX = 1 << 31

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	var dp = make([]int, amount+1)
	for i := range amount + 1 {
		dp[i] = INTMAX
	}

	for _, coin := range coins {
		for i := 0; i <= amount; i++ {
			if i >= coin {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}
	if dp[amount] == INTMAX {
		return -1
	}
	return dp[amount]
}
