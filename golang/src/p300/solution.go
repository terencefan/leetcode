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

func lengthOfLIS_1(nums []int) (r int) {
	if len(nums) == 0 {
		return
	}

	f := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		f[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				f[i] = max(f[i], f[j] + 1)
			}
		}
	}

	for i := 0; i < len(nums); i++ {
		r = max(r, f[i])
	}
	return
}

func binsearch(stack []int, x int) (r int) {
	l, r := 0, len(stack)
	for l < r {
		m := l + (r - l) / 2
		if stack[m] >= x {
			r = m
		} else {
			l = m + 1
		}
	}
	return
}

func lengthOfLIS(nums []int) (r int) {
	stack := make([]int, 0)

	for _, num := range nums {
		r := binsearch(stack, num)
		if r < 0 {
			stack[0] = num
		} else if r == len(stack) {
			stack = append(stack, num)
		} else {
			stack[r] = num
		}
		fmt.Println(stack)
	}

	return len(stack)
}

func main() {
	r := lengthOfLIS([]int{10,9,2,5,3,7,101,18})
	fmt.Println(r)
}
