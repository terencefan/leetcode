package p875

func tryEat(piles []int, k int) int {
	hours := 0
	for _, pile := range piles {
		hours += pile / k
		if pile%k != 0 {
			hours++
		}
	}
	return hours
}

func minEatingSpeed(piles []int, h int) int {
	var max = 0
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}

	l, r := 1, max

	for l < r {
		mid := l + (r-l)/2
		hours := tryEat(piles, mid)
		if hours > h {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
