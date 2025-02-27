package main

import (
	"fmt"
)

const MOD = 1e9 + 7

// 1 2 3
// 4 5 6
// 7 8 9
// * 0 #

var lead = [][]int{
	0: {4, 6},
	1: {6, 8},
	2: {7, 9},
	3: {4, 8},
	4: {0, 3, 9},
	5: {},
	6: {1, 7, 0},
	7: {2, 6},
	8: {1, 3},
	9: {2, 4},
}

func knightDialer(N int) (r int) {
	dp := [][]int{
		{1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}

	for i := 1; i < N; i++ {
		for k := 0; k < 10; k++ {
			next := 0
			for _, pre := range lead[k] {
				next += dp[^i&1][pre]
				next %= MOD
			}
			dp[i&1][k] = next
		}
	}

	for i := 0; i < 10; i++ {
		r += dp[^N&1][i]
		r %= MOD
	}
	fmt.Println(dp)
	return
}

func main() {
	fmt.Println(knightDialer(4))
}
