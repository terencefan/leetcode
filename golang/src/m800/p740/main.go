package p740

import "sort"

func deleteAndEarn(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	prev, hasPrev, noPrev := -1, 0, 0
	for _, num := range nums {
		if num == prev {
			noPrev += num
		} else {
			if num == prev+1 {
				hasPrev, noPrev = max(hasPrev, noPrev), hasPrev+num
			} else {
				hasPrev = max(hasPrev, noPrev)
				noPrev = hasPrev + num
			}
		}
		prev = num
	}
	return max(noPrev, hasPrev)
}
