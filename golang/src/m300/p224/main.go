package p224

func abs(a int, minus bool) int {
	if minus {
		return -a
	}
	return a
}

func calc(s string, i int) (int, int) {
	var minus = false
	var res, val = 0, 0
	for ; i < len(s); i++ {
		c := s[i]

		if c == '-' {
			res += abs(val, minus)
			val = 0
			minus = true
		} else if c == '+' {
			res += abs(val, minus)
			val = 0
			minus = false
		} else if c == '(' {
			val, i = calc(s, i+1)
		} else if c == ')' {
			res += abs(val, minus)
			return res, i + 1
		} else if c >= '0' && c <= '9' {
			val = val*10 + int(c-'0')
		}
	}
	res += abs(val, minus)
	return res, i
}

func calculate(s string) int {
	r, _ := calc(s, 0)
	return r
}
