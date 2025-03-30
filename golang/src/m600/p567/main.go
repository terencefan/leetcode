package p567

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) == 0 {
		return true
	} else if len(s2) == 0 {
		return false
	}

	m, n := make(map[byte]int), make(map[byte]int)

	for i := range len(s1) {
		m[s1[i]]++
	}

	l := 0
	for r := range len(s2) {
		n[s2[r]]++
		for n[s2[r]] > m[s2[r]] && l <= r {
			m[s2[l]]--
			l++
		}
		if r-l == len(s1)-1 {
			return true
		}
	}
	return false
}
