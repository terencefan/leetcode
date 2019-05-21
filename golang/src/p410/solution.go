package main

import (
	"fmt"
)

func iter(nums []int) (s, m int) {
	for _, num := range nums {
		s += num
		if num > m {
			m = num
		}
	}
	return
}

func test(nums []int, k int, m int) bool {
	t := 0
	for _, num := range nums {
		if num > k {
			return false
		} else if t + num > k {
			t = num
			m--
		} else {
			t += num
		}
		if m <= 0 {
			return false
		}
	}
	return true
}


func splitArray(nums []int, m int) int {
	if m > len(nums) {
		return 0
	}

	sum, max := iter(nums)
	if m == 1 {
		return sum
	} else if m == len(nums) {
		return max
	}

	l, r := max, sum
	for l < r {
		mid := l + (r - l) / 2
		if test(nums, mid, m) {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return l
}

func main() {
	// test([]int{7,2,5,10,8}, 14, 2)
	r := splitArray([]int{7,2,5,10,8}, 2)
	fmt.Println(r)
}