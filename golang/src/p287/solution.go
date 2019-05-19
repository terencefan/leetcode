package main

import (
	"fmt"
	"sort"
)

func findDuplicate_1(nums []int) int {
	sort.Ints(nums)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i - 1] {
			return nums[i]
		}
	}
	return -1
}

func findDuplicate_2(nums []int) int {
	b := true
	for b {
		b = false
		for i := 0; i < len(nums) - 1; i++ {
			if nums[i] == nums[i + 1] {
				return nums[i]
			} else if nums[i] > nums[i + 1] {
				nums[i], nums[i + 1] = nums[i + 1], nums[i]
				b = true
			}
		}
	}
	return -1
}

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		j := i
		for j != nums[j] - 1 {
			if nums[nums[j] - 1] == nums[j] {
				return nums[j]
			}
			nums[nums[j] - 1], nums[j] = nums[j], nums[nums[j] - 1]
		}
	}
	return -1
}

func main() {
	r := findDuplicate([]int{8,7,1,10,17,15,18,11,16,9,19,12,5,14,3,4,2,13,18,18})
	fmt.Println(r)
}
