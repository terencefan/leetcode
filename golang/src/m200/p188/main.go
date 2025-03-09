package p188

const INTMIN = -1 << 31

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func maxProfit(k int, prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	buyk := make([]int, k)
	sellk := make([]int, k)

	for i := range k {
		buyk[i] = INTMIN
		sellk[i] = INTMIN
	}

	var r = 0
	for _, price := range prices {
		for j := k - 1; j >= 0; j-- {
			sellk[j] = max(sellk[j], buyk[j]+price)
			r = max(r, sellk[j])
			if j > 0 {
				buyk[j] = max(buyk[j], sellk[j-1]-price)
			} else {
				buyk[0] = max(buyk[0], -price)
			}
		}
	}
	return r
}
