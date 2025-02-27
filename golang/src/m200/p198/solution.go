package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	f, g := make([]int, len(nums)), make([]int, len(nums))
	f[0], g[0] = nums[0], 0
	f[1], g[1] = nums[1], nums[0]

	for i := 2; i < len(nums); i++ {
		f[i] = nums[i] + max(f[i-2], g[i-1])
		g[i] = max(g[i-1], f[i-1])
	}
	return max(f[len(nums)-1], g[len(nums)-1])
}

func main() {
	r := rob([]int{2, 7, 9, 3, 1})
	fmt.Println(r)
}
