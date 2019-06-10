package main

import (
	"fmt"
)

func quickselect(nums []int, lo, hi, k int) int {
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

	if p-lo == k {
		return pivot
	} else if p-lo > k {
		return quickselect(nums, lo, p-1, k)
	} else {
		return quickselect(nums, p+1, hi, k-(p-lo)-1)
	}
}

func count(nums []int, x int, g int) bool {
	for _, num := range nums {
		if num == x {
			g--
		}
		if g == 0 {
			return true
		}
	}
	return false
}

func majorityElement(nums []int) []int {
	if len(nums) < 1 {
		return []int{}
	}
	n := len(nums)

	g := n / 3
	p := quickselect(nums, 0, n-1, g)
	q := quickselect(nums, 0, n-1, n-1-g)

	if p == q {
		if count(nums, p, g) {
			return []int{p}
		} else {
			return []int{}
		}
	} else {
		cp, cq := count(nums, p, g), count(nums, q, g+1)
		if cp && cq {
			return []int{p, q}
		} else if cp {
			return []int{p}
		} else if cq {
			return []int{q}
		} else {
			return []int{}
		}
	}
}

func main() {
	fmt.Println(majorityElement([]int{3, 2, 1, 6, 5, 4, 9, 8, 7}))
}
