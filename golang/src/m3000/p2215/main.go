package p2215

func unique(nums []int) map[int]bool {
	var m = make(map[int]bool)
	for _, num := range nums {
		m[num] = true
	}
	return m
}

func distinct(m1, m2 map[int]bool) []int {
	var r = make([]int, 0)
	for k := range m1 {
		if _, ok := m2[k]; !ok {
			r = append(r, k)
		}
	}
	return r
}

func findDifference(nums1 []int, nums2 []int) [][]int {
	s1, s2 := unique(nums1), unique(nums2)
	return [][]int{
		distinct(s1, s2),
		distinct(s2, s1),
	}
}
