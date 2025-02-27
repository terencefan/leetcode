package main

import (
	"fmt"
)

func parse(expression string, start int) (end int, b bool) {
	N := len(expression)

	if start >= N {
		return N, false
	}
	c := expression[start]

	switch c {
	case 't':
		return start + 1, true
	case 'f':
		return start + 1, false
	case '!':
		if start + 2 >= N {
			return N, false
		}
		if expression[start + 1] != '(' {
			return N, false
		}
		start += 2
		next, r := parse(expression, start)
		b = !r

		if next >= N || expression[next] != ')' {
			return N, false
		} else {
			end = next + 1
			return
		}
	case '|', '&':
		if start + 2 >= N {
			return N, false
		}
		if expression[start + 1] != '(' {
			return N, false
		}
		if c == '&' {
			b = true
		} else {
			b = false
		}

		start += 2

		for start < N {
			next, r := parse(expression, start)
			if c == '&' {
				b = b && r
			} else {
				b = b || r
			}
			if expression[next] == ')' {
				end = next + 1
				return
			} else if expression[next] == ',' {
				start = next + 1
			} else {
				return N, false
			}
		}

	}
	return
}

func parseBoolExpr(expression string) bool {
	if len(expression) == 0 {
		return false
	}
	_, b := parse(expression, 0)
	return b
}

func main() {
	r := parseBoolExpr("&(|(t,f,f),!(f))")
	fmt.Println(r)
}
