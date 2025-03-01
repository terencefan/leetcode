package p1732

func largestAltitude(gain []int) int {
	var c, r = 0, 0
	for _, v := range gain {
		c += v
		if c > r {
			r = c
		}
	}
	return r
}
