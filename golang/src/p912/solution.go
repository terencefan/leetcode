package main

import (
	"sort"
)

func quicksort(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}

	l, r := lo, hi
	p, pivot := lo, nums[lo]

	for l < r {
		for ; l < r; r-- {
			if nums[r] < pivot {
				nums[p] = nums[r]
				p = r
				break
			}
		}
		for ; l < r; l++ {
			if nums[l] > pivot {
				nums[p] = nums[l]
				p = l
				break
			}
		}
	}
	nums[p] = pivot

	quicksort(nums, lo, p-1)
	quicksort(nums, p+1, hi)
}

func sortArray(nums []int) []int {
	sort.Ints(nums)
	return nums
}
