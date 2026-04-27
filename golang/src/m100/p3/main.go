package p3

func lengthOfLongestSubstring(s string) int {
	var m = make(map[rune]bool)
	var l, r = 0, 0

	for _, c := range s {
		if _, ok := m[c]; !ok {
			m[c] = true
			r = max(r, len(m))
			continue
		}

		for rune(s[l]) != c {
			delete(m, rune(s[l]))
			l++
		}
		l++ // skip c
	}
	return r
}
