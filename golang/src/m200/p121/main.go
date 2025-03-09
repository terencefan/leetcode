package p121

const INTMIN = -1 << 31

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxProfit(prices []int) int {
	buy, sell := INTMIN, INTMIN
	for _, price := range prices {
		buy = max(buy, -price)
		sell = max(sell, buy+price)
	}
	return sell
}
