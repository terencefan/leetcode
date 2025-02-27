package main

func intersect(nums1 []int, nums2 []int) (r []int) {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	var m = make(map[int]int)
	for _, num := range nums1 {
		m[num]++
	}
	for _, num := range nums2 {
		if m[num] > 0 {
			r = append(r, num)
			m[num]--
		}
	}
	return
}
