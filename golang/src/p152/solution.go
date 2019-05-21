package main

import (
	"fmt"
	"math"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	f, g := make([]int, len(nums)), make([]int, len(nums))
	f[0] = nums[0]
	g[0] = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			continue
		} else if nums[i] > 0 {
			f[i] = max(f[i - 1] * nums[i], nums[i])
			g[i] = min(g[i - 1] * nums[i], nums[i])
		} else {
			f[i] = max(g[i - 1] * nums[i], nums[i])
			g[i] = min(f[i - 1] * nums[i], nums[i])
		}
	}

	r := -math.MaxInt32
	for i := 0; i < len(nums); i++ {
		r = max(r, f[i])
	}
	return r
}

func main() {
	r := maxProduct([]int{-2, 1, -9})
	fmt.Println(r)
}
