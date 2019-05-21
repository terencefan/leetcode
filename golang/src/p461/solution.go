package main

func hammingDistance(x int, y int) int {
	var z int
	z = x ^ y
	var n = 0
	for z > 0 {
		z = z & (z - 1)
		n++
	}
	return n
}
