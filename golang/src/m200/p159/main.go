package p159

func lengthOfLongestSubstringTwoDistinct(s string) int {
	l := 0
	m := make(map[byte]int)

	res := 0
	for r := range len(s) {
		rc := s[r]
		m[rc]++

		for len(m) > 2 {
			lc := s[l]
			m[lc]--
			if m[lc] == 0 {
				delete(m, lc)
			}
			l++
		}
		res = max(res, r-l+1)
	}
	return res
}
