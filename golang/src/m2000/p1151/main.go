package p1151

func minSwaps(data []int) int {
	count := 0

	for _, num := range data {
		if num == 1 {
			count++
		}
	}

	i, swaps := 0, 0
	for ; i < count; i++ {
		if data[i] == 0 {
			swaps++
		}
	}

	var r = swaps
	for ; i < len(data); i++ {
		nl, nr := data[i-count], data[i]
		if nl == 0 && nr == 1 {
			swaps--
			r = min(r, swaps)
		} else if nl == 1 && nr == 0 {
			swaps++
		}
	}
	return r
}
