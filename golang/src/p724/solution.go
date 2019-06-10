package main

func pivotIndex(nums []int) int {

	right := 0
	for _, num := range nums {
		right += num
	}

	left := 0
	for i := 0; i < len(nums); i++ {
		right -= nums[i]
		if left == right {
			return i
		}
		left += nums[i]
	}
	return -1
}
