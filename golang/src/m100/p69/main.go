package p69

func mySqrt(x int) int {
	if x == 1 {
		return 1
	}

	i, j := 0, x / 2 + 1

	for i < j {
		mid := i + (j - i) / 2

		s := mid * mid
		if s == x {
			return mid
		} else if s > x {
			j = mid
		} else {
			i = mid + 1
		}
	}
	return i
}