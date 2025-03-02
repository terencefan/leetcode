package p1207

func uniqueOccurrences(arr []int) bool {
	var m = make(map[int]int)
	for _, v := range arr {
		m[v]++
	}

	var n = make(map[int]bool)
	for _, count := range m {
		if _, ok := n[count]; ok {
			return false
		}
		n[count] = true
	}
	return true
}
