package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	} else if len(cost) == 1 {
		return cost[1]
	}

	p, q := 0, 0
	for i := 2; i <= len(cost); i++ {
		p, q = q, min(cost[i-2]+p, cost[i-1]+q)
	}
	return q
}

func main() {
	r := minCostClimbingStairs([]int{1, 0, 0, 0})
	fmt.Println(r)
}
