package main

import (
	"fmt"
)

func majorityElement(nums []int) int {
	m := make(map[int]int)
	for _, num := range nums {
		m[num]++
		if m[num] > len(nums) / 2 {
			return num
		}
	}
	return -1
}

func main() {
	r := majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	fmt.Println(r)
}
