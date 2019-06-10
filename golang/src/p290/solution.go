package main

func wordPattern(pattern string, str string) bool {
	var m = make(map[byte]string)
	var n = make(map[string]byte)

	l, r := 0, 0
	for i := range pattern {
		for r < len(str) && str[r] != ' ' {
			r++
		}
		if l == r {
			return false
		}

		segment := str[l:r]
		if s, ok := m[pattern[i]]; ok {
			if s != segment {
				return false
			}
		} else {
			m[pattern[i]] = segment
			if n[segment] != 0 {
				return false
			}
			n[segment] = pattern[i]
		}
		r += 1
		l = r
	}
	return r == len(str)+1
}
