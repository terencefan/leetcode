package p88

func merge(nums1 []int, m int, nums2 []int, n int) {
	l, r := m-1, n-1

	for l >= 0 && r >= 0 {
		if nums1[l] > nums2[r] {
			nums1[l+r+1] = nums1[l]
			l--
		} else {
			nums1[l+r+1] = nums2[r]
			r--
		}
	}

	for ; r >= 0; r-- {
		nums1[r] = nums2[r]
	}
}
