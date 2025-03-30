package p1011

func canShipInDays(weights []int, days, capacity int) bool {
	current := 0
	for _, weight := range weights {
		if current > capacity {
			return false
		} else if current+weight > capacity {
			days--
			current = weight
		} else {
			current += weight
		}
		if days < 0 {
			return false
		}
	}
	return current == 0 || days > 0
}

func shipWithinDays(weights []int, days int) int {
	s := 0
	for _, weight := range weights {
		s += weight
	}

	l, r := 1, s

	for l < r {
		mid := l + (r-l)/2

		if canShipInDays(weights, days, mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}
