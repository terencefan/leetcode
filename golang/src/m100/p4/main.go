package p4

const INTMAX = 1 << 31
const INTMIN = -1 << 31

func get(nums []int, i int) int {
	if i < 0 {
		return INTMIN
	} else if i >= len(nums) {
		return INTMAX
	}
	return nums[i]
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	m, n := len(nums1), len(nums2)

	var left1, left2, right1, right2 = 0, 0, 0, 0

	// [1,3,5]
	// [1,4,7]

	l, r := 0, m
	for l <= r {
		mid1 := l + (r-l)/2
		mid2 := (m+n+1)/2 - mid1

		left1 = get(nums1, mid1-1)
		right1 = get(nums1, mid1)
		left2 = get(nums2, mid2-1)
		right2 = get(nums2, mid2)

		if right1 >= left2 && right2 >= left1 {
			if (m+n)%2 == 0 {
				return float64(max(left1, left2)+min(right1, right2)) / 2.0
			} else {
				return float64(max(left1, left2))
			}
		} else if right1 < left2 {
			l = mid1 + 1
		} else if left1 > right2 {
			r = mid1 - 1
		}
	}
	return 0
}
