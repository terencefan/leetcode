package main

import (
	"fmt"
)

func hash(num, t int) int {
	if num >= 0 {
		return num / t
	} else {
		return (num+1)/t - 1
	}
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	if k < 1 || t < 0 {
		return false
	}

	var m = make(map[int]int)

	for i, num := range nums {
		h := hash(num, t+1)

		if _, ok := m[h]; ok {
			return true
		}
		if v, ok := m[h-1]; ok && num-v <= t {
			return true
		}
		if v, ok := m[h+1]; ok && v-num <= t {
			return true
		}
		m[h] = num

		if i >= k {
			index := hash(nums[i-k], t+1)
			delete(m, index)
		}
	}
	return false
}

func main() {
	r := containsNearbyAlmostDuplicate([]int{2, 0, -2, 2}, 2, 1)
	fmt.Println(r)
}
