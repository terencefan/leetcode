package p191

func hammingWeight(n int) int {
	var c, x = 0, 1

	for x <= n {
		if x&n != 0 {
			c++
		}
		if x == 1<<31 {
			break
		}
		x <<= 1
	}
	return c
}
