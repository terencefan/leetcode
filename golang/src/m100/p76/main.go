package p76

func minWindow(s string, t string) string {

	var target = make(map[byte]int)
	var current = make(map[byte]int)

	for i := range t {
		target[t[i]]++
	}
	var ck, tk = 0, len(target)

	var res string
	l, r := -1, -1
	for i := range s {
		c := s[i]
		if target[c] > 0 {
			if l < 0 {
				l = i
			}
			current[c]++
			if current[c] == target[c] {
				ck++
			}
		}
		if ck == tk {
			r = i
			for l <= r {
				c := s[l]
				l++
				if target[c] > 0 && current[c] == target[c] {
					if len(res) == 0 || r-l+2 < len(res) {
						res = s[l-1 : r+1]
					}
					current[c]--
					ck--
					break
				} else if current[c] > 0 {
					current[c]--
				}
			}
		}
	}
	return res
}
