package main

import (
	"fmt"
)

func quickfind(nums []int, k int) int {
	i, j := 0, len(nums) - 1

	pivot, p := nums[0], 0
	for i <= j {
		for ;j>=p;j-- {
			if nums[j] > pivot {
				nums[p] = nums[j]
				p = j
				break
			}
		}

		for ;i<=p;i++ {
			if nums[i] < pivot {
				nums[p] = nums[i]
				p = i
				break
			}
		}
	}

	if p == k - 1 {
		return pivot
	} else if p < k - 1 {
		return quickfind(nums[p+1:], k - p - 1)
	} else {
		return quickfind(nums[:p], k)
	}
}

func findKthLargest(nums []int, k int) int {
	return quickfind(nums, k)
}

func main() {
	r := findKthLargest([]int{2,1,2,4}, 2)
	fmt.Println(r)
}
