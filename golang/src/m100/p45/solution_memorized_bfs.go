package main

import (
	"fmt"
)

func minp(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func jump(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	q, dp := []int{0}, make([]int, len(nums))

	steps := 0
	for len(q) > 0 {
		length := len(q)

		for i := 0; i < length; i++ {
			pos := q[i]

			for k := minp(nums[pos], len(nums)-pos-1); k >= 1; k-- {
				if pos+k == len(nums)-1 {
					return steps + 1
				} else if dp[pos+k] > 0 && dp[pos+k] <= steps+1 {
					continue
				} else {
					dp[pos+k] = steps + 1
					q = append(q, pos+k)
				}
			}
		}
		q = q[length:]
		steps++
	}
	return 0
}

func main() {
	r := jump([]int{2, 3, 1, 1, 4})
	fmt.Println(r)
}
