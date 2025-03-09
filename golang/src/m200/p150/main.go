package p150

func stoi(s string) int {
	r, minus := 0, false
	for _, c := range s {
		switch c {
		case '-':
			minus = true
		default:
			r = r*10 + int(c-'0')
		}
	}
	if minus {
		return -r
	} else {
		return r
	}
}

func calc(a, b int, op byte) int {
	switch op {
	case '+':
		return a + b
	case '-':
		return a - b
	case '*':
		return a * b
	case '/':
		return a / b
	}
	return 0
}

func evalRPN(tokens []string) int {
	s := make([]int, 0)

	for _, token := range tokens {
		if len(token) == 1 {
			switch token[0] {
			case '+', '-', '*', '/':
				if len(s) < 2 {
					return -1
				}
				op2 := s[len(s)-1]
				op1 := s[len(s)-2]
				s[len(s)-2] = calc(op1, op2, token[0])
				s = s[:len(s)-1]
			default:
				s = append(s, stoi(token))
			}
		} else {
			s = append(s, stoi(token))
		}
	}
	if len(s) != 1 {
		return -1
	}
	return s[0]
}
