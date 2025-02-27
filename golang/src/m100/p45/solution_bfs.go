package main

import (
	"fmt"
)

func jump_bfs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	q := []int{0}

	steps := 0
	for len(q) > 0 {
		length := len(q)
		m := make(map[int]bool)

		for i := 0; i < length; i++ {
			pos := q[i]

			for k := nums[pos]; k >= 1; k-- {
				if pos+k > len(nums)-1 {
					continue
				} else if pos+k == len(nums)-1 {
					return steps + 1
				} else {
					m[pos+k] = true
				}
			}
		}

		q = q[length:]
		for pos := range m {
			q = append(q, pos)
		}
		steps++
	}
	return 0
}

func main() {
	r := jump_bfs([]int{2, 3, 1, 1, 4})
	fmt.Println(r)
}
