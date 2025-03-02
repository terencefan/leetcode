package p714

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}

	n := len(prices)

	free, hold := 0, -prices[0]

	for i := 1; i < n; i++ {
		t := free
		free = max(free, hold+prices[i]-fee)
		hold = max(hold, t-prices[i])
	}

	return free
}
