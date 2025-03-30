package p678

func checkValidString(s string) bool {
	var p, r = make([]int, 0), make([]int, 0)

	for i, c := range s {
		switch c {
		case '(':
			p = append(p, i)
		case ')':
			if len(p) > 0 {
				p = p[:len(p)-1]
			} else if len(r) > 0 {
				r = r[:len(r)-1]
			} else {
				return false
			}
		case '*':
			r = append(r, i)
		}
	}

	for len(p) > 0 && len(r) > 0 {
		if p[len(p)-1] > r[len(r)-1] {
			return false
		}
		p = p[:len(p)-1]
		r = r[:len(r)-1]
	}
	return len(p) <= len(r)
}
