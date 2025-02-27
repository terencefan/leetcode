package main

import (
	"fmt"
)

func findPeakElement(nums []int) int {
	l, r := 0, len(nums)-1

	for l < r {
		m := l + (r-l)/2
		if nums[m] > nums[m+1] {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}

func main() {
	fmt.Println(findPeakElement([]int{1, 3, 2, 1}))
}
