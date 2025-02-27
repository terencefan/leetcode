package main

func singleNumber(nums []int) int {
	var r = 0
	for _, n := range nums {
		r ^= n
	}
	return r
}
