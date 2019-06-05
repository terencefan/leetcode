package main

import (
	"fmt"
)

func firstMissingPositive(nums []int) int {
	if len(nums) == 0 {
		return 1
	}

	for i := range nums {
		num := nums[i]
		for num > 0 && num <= len(nums) && num != nums[num-1] {
			nums[i], nums[num - 1] = nums[num - 1], nums[i]
			num = nums[i]
		}
	}

	for i, num := range nums {
		if i + 1 != num {
			return i + 1
		}
	}
	return len(nums) + 1
}

func main() {
	r := firstMissingPositive([]int{3})
	fmt.Println(r)
}
