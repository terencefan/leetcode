package main

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	l := len(nums) + 1

	i, j, sum := 0, 0, nums[0]
	for {
		for sum < s && j < len(nums)-1 {
			j++
			sum += nums[j]
		}
		if sum < s {
			break
		}
		for sum-nums[i] >= s && i < j {
			sum -= nums[i]
			i++
		}
		l = min(l, j-i+1)
		sum -= nums[i]
		i++
	}

	if l == len(nums)+1 {
		return 0
	} else {
		return l
	}
}
