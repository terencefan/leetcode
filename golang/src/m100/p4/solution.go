package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func findMedian(nums []int) float64 {
	if len(nums)%2 == 0 {
		return float64(nums[len(nums)/2-1])/2 + float64(nums[len(nums)/2])/2
	} else {
		return float64(nums[len(nums)/2])
	}
}

func countLessThan(nums []int, t float64) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := lo + (hi-lo)/2
		if float64(nums[mid]) < t {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return lo
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 && len(nums2) == 0 {
		return 0.0
	} else if len(nums1) == 0 {
		return findMedian(nums2)
	} else if len(nums2) == 0 {
		return findMedian(nums1)
	}

	length := len(nums1) + len(nums2)
	pos := length / 2
	lo, hi := float64(min(nums1[0], nums2[0])), float64(max(nums1[len(nums1)-1], nums2[len(nums2)-1]))

	for {
		mid := lo + (hi-lo)/2

		p := countLessThan(nums1, mid)
		q := countLessThan(nums2, mid)

		if p+q == pos {
			if length%2 == 1 {
				if p == len(nums1) {
					return float64(nums2[q])
				} else if q == len(nums2) {
					return float64(nums1[p])
				} else {
					return float64(min(nums1[p], nums2[q]))
				}
			} else {
				var a, b int
				if p == len(nums1) {
					b = nums2[q]
				} else if q == len(nums2) {
					b = nums1[p]
				} else {
					b = min(nums1[p], nums2[q])
				}

				if p == 0 {
					a = nums2[q-1]
				} else if q == 0 {
					a = nums1[p-1]
				} else {
					a = max(nums1[p-1], nums2[q-1])
				}
				return float64(a+b) / 2
			}
		} else if p+q > pos {
			hi = mid
		} else {
			lo = mid + 0.1
		}

		if hi-lo < 0.2 {
			if p == 0 {
				return float64(nums2[q])
			} else if p == len(nums1) {
				return float64(nums2[q-1])
			} else if q == 0 {
				return float64(nums1[p])
			} else if q == len(nums2) {
				return float64(nums1[p-1])
			} else {
				return float64(min(nums1[p], nums2[q]))
			}
		}
	}
	return 0.0
}

func main() {
	r := findMedianSortedArrays([]int{1, 1}, []int{1, 2})
	fmt.Println(r)
}
