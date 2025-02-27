package main

import (
	"fmt"
)

var cache map[int]bool

func dfs(nums []int, state, partitionSum, currentSum, k int) bool {
	if k == 0 {
		return true
	}

	if b, ok := cache[state]; ok {
		return b
	}

	s := 1
	for i := 0; i < len(nums); i++ {
		if state&s != 0 {
			// do nothing
		} else {
			if nums[i] > partitionSum-currentSum {
				// do nothing.
			} else if nums[i] == partitionSum-currentSum {
				if dfs(nums, state|s, partitionSum, 0, k-1) {
					cache[state] = true
					return true
				}
			} else {
				if dfs(nums, state|s, partitionSum, currentSum+nums[i], k) {
					cache[state] = true
					return true
				}
			}
		}
		s <<= 1
	}

	cache[state] = false
	return false
}

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%k != 0 {
		return false
	}
	partitionSum := sum / k

	cache = make(map[int]bool)
	return dfs(nums, 0, partitionSum, 0, k)
}

func main() {
	fmt.Println(canPartitionKSubsets([]int{4, 3, 2, 3, 5, 2, 1}, 4))
}
