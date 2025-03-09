package p15

import "slices"

const INTMIN = -1 << 31

func find(nums []int, index, target int) [][]int {
	prev, l, r := INTMIN, index, len(nums)-1

	res := make([][]int, 0)
	for l < r {
		sum := nums[l] + nums[r]
		if nums[l] == prev {
			prev = nums[l]
			l++
		} else if sum == target {
			res = append(res, []int{nums[l], nums[r]})
			prev = nums[l]
			l++
		} else if sum < target {
			prev = nums[l]
			l++
		} else {
			r--
		}
	}
	return res
}

func threeSum(nums []int) [][]int {
	slices.Sort(nums)

	var r = make([][]int, 0)

	var prev = INTMIN
	for i, v := range nums {
		if v != prev {
			for _, pair := range find(nums, i+1, -v) {
				r = append(r, []int{v, pair[0], pair[1]})
			}
		}
		prev = v
	}
	return r
}
