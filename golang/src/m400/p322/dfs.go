package p322

import "sort"

func dfs(coins []int, cache []int, amount, count int) {
	if cache[amount] != 0 && cache[amount] <= count {
		return
	}
	cache[amount] = count

	for _, coin := range coins {
		if amount >= coin {
			dfs(coins, cache, amount-coin, count+1)
		}
	}
}

func coinChange2(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	var cache = make([]int, amount+1)

	sort.Slice(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	dfs(coins, cache, amount, 0)
	if cache[0] == 0 {
		return -1
	}
	return cache[0]
}
