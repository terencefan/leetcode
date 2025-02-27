package main

import "math/bits"

func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

func minFlips2(a int, b int, c int) int {
	if a == 0 && b == 0 && c == 0 {
		return 0
	}

	x, y, z := a%2, b%2, c%2

	r := minFlips2(a>>1, b>>1, c>>1)
	if z == 0 {
		return r + btoi(x == 1) + btoi(y == 1)
	} else {
		return r + btoi(x == 0 && y == 0)
	}
}

func minFlips(a int, b int, c int) int {
	var x, y, z uint
	x, y, z = uint(a), uint(b), uint(c)
	diffs := (x | y) ^ z
	return bits.OnesCount(diffs) + bits.OnesCount(x&y&diffs)
}
