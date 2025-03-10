package p3

func lengthOfLongestSubstring(s string) int {
	var m = make(map[byte]bool)
	var l, r = 0, 0

	for i := range s {
		c := s[i]
		if _, ok := m[c]; !ok {
			m[c] = true
			r = max(r, len(m))
			continue
		}

		for s[l] != c {
			delete(m, s[l])
			l++
		}
		l++
	}
	return r
}
