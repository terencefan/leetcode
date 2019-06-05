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

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}

	currentMin, currentMax, steps := 0, 0, 0
	for currentMax < len(nums) - 1 {
		temp := currentMax
		for i := currentMin; i <= temp && i < len(nums); i++ {
			currentMax = max(currentMax, i + nums[i])
		}
		currentMin = temp + 1
		steps++
	}
	return steps
}

func main() {
	r := jump([]int{2,3,1,1,4})
	fmt.Println(r)
}
