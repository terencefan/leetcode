package p16

import "slices"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	slices.Sort(nums)

	res := nums[0] + nums[1] + nums[2]
	for i := range len(nums) - 2 { // Fix loop range
		// Skip duplicate values for the first number
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			if sum == target {
				return target
			}
			if abs(sum-target) < abs(res-target) {
				res = sum
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
