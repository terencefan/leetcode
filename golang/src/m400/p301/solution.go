package p301

const INTMIN = -1 << 31
const INTMAX = 1 << 31

func recursive(m map[string]bool, s string, bytes []byte, lc, rc, clc, idx int) {
	if idx == len(s) {
		if lc == rc {
			m[string(bytes)] = true
		}
		return
	}

	b := s[idx]

	if b == '(' {
		if lc > 0 {
			// remove this '('
			recursive(m, s, bytes, lc-1, rc, clc, idx+1)
		}
		recursive(m, s, append(bytes, b), lc, rc, clc+1, idx+1)
	} else if b == ')' {
		if rc > 0 {
			// remove this ')'
			recursive(m, s, bytes, lc, rc-1, clc, idx+1)
		}
		if clc > 0 {
			recursive(m, s, append(bytes, b), lc, rc, clc-1, idx+1)
		}
	} else {
		recursive(m, s, append(bytes, b), lc, rc, clc, idx+1)
	}
}

func removeInvalidParentheses(s string) []string {
	x, minx := 0, INTMAX
	for _, c := range s {
		if c == '(' {
			x++
		} else if c == ')' {
			x--
		}
		minx = min(x, minx)
	}

	var bytes = make([]byte, 0)
	var m = make(map[string]bool, 0)

	if minx < 0 {
		recursive(m, s, bytes, x-minx, max(-x, -minx), 0, 0)
	} else {
		recursive(m, s, bytes, x, 0, 0, 0)
	}

	keys := make([]string, 0)
	for key := range m {
		keys = append(keys, key)
	}
	return keys
}
