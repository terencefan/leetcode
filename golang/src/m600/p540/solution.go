package main

func singleNonDuplicate(nums []int) int {
	var i, j, k = 0, len(nums)/2, 0
	for i < j {
		k = (i + j) / 2
		if nums[2*k] == nums[2*k+1] {
			i = k + 1
		} else {
			j = k
		}
	}
	return nums[2*i]
}
