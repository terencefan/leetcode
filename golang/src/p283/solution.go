package main

import (
	"fmt"
)

func moveZeroes(nums []int) {
	offset := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			offset++
		} else if offset > 0 {
			nums[i], nums[i-offset] = nums[i-offset], nums[i]
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums)
}
