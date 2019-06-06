package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}

	dp := [][]int{
		{costs[0][0], costs[0][1], costs[0][2]},
		{0, 0, 0},
	}

	for i := 1; i < len(costs); i++ {
		p := (i + 1) % 2
		dp[i%2][0] = min(dp[p][1], dp[p][2]) + costs[i][0]
		dp[i%2][1] = min(dp[p][0], dp[p][2]) + costs[i][1]
		dp[i%2][2] = min(dp[p][0], dp[p][1]) + costs[i][2]
	}
	p := (len(costs) - 1) % 2
	return min(dp[p][0], min(dp[p][1], dp[p][2]))
}

func main() {
	costs := [][]int{
		{7, 2, 6},
		{13, 13, 5},
		{16, 3, 1},
		{2, 6, 19},
		{4, 8, 16},
		{20, 10, 17},
		{4, 2, 1},
		{9, 6, 1},
		{15, 6, 4},
		{3, 8, 5},
		{12, 14, 7},
		{8, 7, 1},
		{19, 12, 8},
		{17, 16, 13},
		{4, 3, 13},
		{8, 5, 12},
		{14, 3, 3},
		{1, 4, 19},
		{7, 4, 7},
		{3, 6, 11},
		{7, 16, 9},
		{20, 4, 1},
		{15, 8, 16},
		{19, 6, 8},
		{19, 12, 3},
		{12, 8, 5},
		{12, 18, 11},
		{8, 11, 15},
		{13, 5, 4},
		{16, 3, 15},
		{13, 16, 7},
		{17, 10, 12},
		{3, 12, 12},
		{15, 15, 18},
		{18, 6, 6},
		{17, 13, 11},
		{7, 19, 19},
		{18, 18, 12},
		{7, 20, 4},
	}
	fmt.Println(minCost(costs))
}
