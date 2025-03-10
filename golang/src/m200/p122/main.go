package p122

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var buy, sell = -prices[0], 0
	for _, price := range prices {
		buy = max(buy, sell-price)
		sell = max(sell, buy+price)
	}
	return sell
}
