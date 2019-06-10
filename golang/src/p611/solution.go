package main

import (
	"fmt"
)

func quicksort(nums []int, l, r int) {
	if l >= r {
		return
	}

	lo, hi := l, r
	p, pivot := lo, nums[lo]

	for lo < hi {
		for ; lo < hi; hi-- {
			if nums[hi] < pivot {
				nums[p] = nums[hi]
				p = hi
				break
			}
		}
		for ; lo < hi; lo++ {
			if nums[lo] > pivot {
				nums[p] = nums[lo]
				p = lo
				break
			}
		}
	}
	nums[p] = pivot

	quicksort(nums, l, p-1)
	quicksort(nums, p+1, r)
}

func find(nums []int, x int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if nums[mid] < x {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	quicksort(nums, 0, len(nums)-1)

	r := 0
	for i := 0; i < len(nums)-2; i++ {
		j := i + 1
		k := find(nums, nums[i]+nums[j])
		for ; j < len(nums)-1; j++ {
			for ; k < len(nums) && nums[i]+nums[j] > nums[k]; k++ {
			}
			if k-j-1 > 0 {
				r += k - j - 1
			}
		}
	}
	return r
}

func main() {
	r := triangleNumber([]int{2, 2, 3, 4})
	fmt.Println(r)
}
