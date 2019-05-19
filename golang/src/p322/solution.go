package main

import (
	"fmt"
)

type Node struct {
	count, delta int
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func coinChange(coins []int, amount int) int {
	f := make([]int, amount + 1)
	for i := 1; i <= amount; i++ {
		f[i] = amount + 1
	}

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			f[j] = min(f[j], f[j - coins[i]] + 1)
		}
	}
	if f[amount] > amount {
		return -1
	} else {
		return f[amount]
	}
}

func main() {
	r := coinChange([]int{1,2,5}, 100)
	fmt.Println(r)
}
