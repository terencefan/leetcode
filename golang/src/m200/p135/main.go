package p135

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func candy(ratings []int) int {
	var candy = make([]int, len(ratings))

	candy[0] = 1
	for i := 1; i < len(ratings); i++ {
		if ratings[i] > ratings[i-1] {
			candy[i] = candy[i-1] + 1
		} else {
			candy[i] = 1
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			candy[i] = max(candy[i], candy[i+1]+1)
		}
	}

	r := 0
	for _, v := range candy {
		r += v
	}
	return r
}
