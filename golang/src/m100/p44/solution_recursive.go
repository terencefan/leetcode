package main

var m [][]int

func isMatch1(s string, p string) bool {
	m = make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		m[i] = make([]int, len(p))
	}

	if len(p) == 0 {
		return len(s) == 0
	}
	r := []byte{p[0]}
	for i := 1; i < len(p); i++ {
		if r[len(r)-1] == '*' {
			if p[i] == '*' {
				continue
			}
		}
		r = append(r, p[i])
	}
	return match(s, string(r), 0, 0)
}

func match(s, p string, i, j int) bool {
	if j == len(p)-1 && p[j] == '*' {
		if i < len(s) {
			m[i][j] = 1
		}
		return true
	} else if i == len(s) || j == len(p) {
		return i == len(s) && j == len(p)
	} else if m[i][j] > 0 {
		return m[i][j] == 1
	}

	if p[j] == '?' {
		return match(s, p, i+1, j+1)
	} else if p[j] == '*' {
		pnext := p[j+1]
		for k := i; k < len(s); k++ {
			if s[k] == pnext || pnext == '?' {
				if match(s, p, k, j+1) {
					m[i][j] = 1
					return true
				}
			}
		}
		m[i][j] = 2
		return false

	} else if s[i] != p[j] {
		m[i][j] = 2
		return false
	} else {
		return match(s, p, i+1, j+1)
	}
}
