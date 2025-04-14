package p16

import "sort"

func threeSumClosest_sort(nums []int, target int) int {
	n := len(nums)
	if n < 3 {
		return 0
	}

	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	var abs = func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	var find = func(a, b int) int {
		k1, k2 := nums[a], nums[b]
		s := k1 + k2

		l, r := a+1, b-1
		for l < r {
			m := l + (r-l)/2
			if s+nums[m] < target {
				l = m + 1
			} else {
				r = m
			}
		}

		res := s + nums[l]
		if l-1 > a && abs(s+nums[l-1]-target) < abs(res-target) {
			res = s + nums[l-1]
		}
		return res
	}

	var res = 1 << 31
	for i1 := range nums {
		for i2 := i1 + 2; i2 < len(nums); i2++ {
			t := find(i1, i2)
			if abs(t-target) < abs(res-target) {
				res = t
			}
		}
	}
	return res
}
