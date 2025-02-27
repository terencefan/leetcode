package main

import (
	"fmt"
)

func subarraySum(nums []int, k int) int {
	hmap := make(map[int]int)
	hmap[0] = 1

	r, p := 0, 0
	for i, num := range nums {
		_ = i
		p += num
		r += hmap[p-k]
		hmap[p]++
	}
	return r
}

func main() {
	r := subarraySum([]int{1, 2, 1}, 2)
	fmt.Println(r)
}
