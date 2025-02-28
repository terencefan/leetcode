package p1679

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func maxOperations(nums []int, k int) int {
	var m1 = make(map[int]int)
	var m2 = make(map[int]int)

	for _, num := range nums {
		if num > k {
			continue
		} else if num > k/2 {
			m2[num]++
		} else {
			m1[num]++
		}
	}

	r := 0

	if k%2 == 0 {
		r += m1[k/2] / 2
	}

	for num, c2 := range m2 {
		if c1, ok := m1[k-num]; ok {
			r += min(c1, c2)
		}
	}
	return r
}
