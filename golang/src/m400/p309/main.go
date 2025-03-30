package p309

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	hold, sell, cooldown := -prices[0], 0, 0

	for i := 1; i < len(prices); i++ {
		prevHold, prevSell := hold, sell
		hold = max(hold, cooldown-prices[i])
		sell = max(sell, prevHold+prices[i])
		cooldown = max(cooldown, prevSell)
	}
	return max(sell, cooldown)
}
