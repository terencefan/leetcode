package p123

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	n := len(prices)

	buy1 := make([]int, n)
	sell1 := make([]int, n)
	buy2 := make([]int, n)
	sell2 := make([]int, n)

	buy1[0] = -prices[0]
	sell1[0] = -1 << 31
	buy2[0] = -1 << 31
	sell2[0] = -1 << 31

	r := 0
	for i := 1; i < n; i++ {
		buy1[i] = max(buy1[i-1], -prices[i])
		sell1[i] = max(sell1[i-1], buy1[i-1]+prices[i])
		r = max(r, sell1[i])
		buy2[i] = max(buy2[i-1], sell1[i-1]-prices[i])
		sell2[i] = max(sell2[i-1], buy2[i-1]+prices[i])
		r = max(r, sell2[i])
	}
	return r
}
