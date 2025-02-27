package main

import (
	"fmt"
)

func canCross_n3(stones []int) bool {

	N := len(stones)

	var f [1100][1100]bool

	if N == 0 {
		return false
	} else if N == 1 {
		return stones[0] == 0
	}

	f[0][1] = stones[0] == 0 && stones[1] == 1

	for i := 2; i < N; i++ {
		for j := 1; j < i; j++ {

			delta := stones[i] - stones[j]

			lo, hi := 0, j
			for lo < hi {
				mid := lo + (hi - lo) / 2
				if stones[j] - stones[mid] > delta + 1 {
					lo = mid + 1
				} else {
					hi = mid
				}
			}

			for k := lo; k < j && stones[j] - stones[k] >= delta - 1; k++ {
				if f[k][j] {
					f[j][i] = true
					break
				}
			}
		}
	}

	for i := 0; i < N; i++ {
		if f[i][N - 1] {
			return true
		}
	}
	return false
}

func main() {
	r := canCross([]int{0,1,2,3,4,8,9,11})
	fmt.Println(r)
}
