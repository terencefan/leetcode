package main

import (
	"fmt"
)

const (
	BGN = '$'
	ADD = '+'
	SUB = '-'
	MUL = '*'
	DIV = '/'
	LPT = '('
	RPT = ')'
)

var priorities = map[int]int {
	BGN: 0,
	ADD: 2,
	SUB: 2,
	MUL: 3,
	DIV: 3,
}

type Function func(a, b int) int

var functions = map[int]Function{
	ADD: func(a, b int) int {
		return a + b
	},
	SUB: func(a, b int) int {
		return a - b
	},
	MUL: func(a, b int) int {
		return a * b
	},
	DIV: func(a, b int) int {
		return a / b
	},
}

type Stack []int

func (s Stack) len() int {
	return len(s)
}

func (s Stack) peek() int {
	return s[s.len() - 1]
}

func (s *Stack) push(x int) {
	*s = append(*s, x)
}

func (s *Stack) pop() (r int) {
	r = s.peek()
	*s = (*s)[:s.len() - 1]
	return
}

func round(nums *Stack, opt int) (num int) {
	function := functions[opt]
	a, b := nums.pop(), nums.pop()
	num = function(b, a)
	nums.push(num)
	return
}

func calculate(s string) int {
	nums := make(Stack, 0)
	opts := make(Stack, 0)
	opts.push('$')

	num := 0
	for i := 0; i < len(s); i++ {
		c := int(s[i])

		if c == RPT {
			nums.push(num)
			for {
				opt := opts.pop()
				if opt == LPT {
					break
				}
				num = round(&nums, opt)
			}
			num = nums.pop()
		} else if c == LPT {
			opts.push(c)
		} else if p, ok := priorities[c]; ok {
			for priorities[opts.peek()] >= p {
				num = functions[opts.pop()](nums.pop(), num)
			}
			nums.push(num)
			opts.push(c)
			num = 0
		} else if c >= '0' && c <= '9' {
			num *= 10
			num += int(c - '0')
		} else if c == ' ' {
			continue
		} else {
			return 0
		}
	}
	nums.push(num)

	for opts.peek() != BGN {
		round(&nums, opts.pop())
	}
	return nums.peek()
}

func main() {
	r := calculate("5 * (1 + 2) + (5 * 3 - 2)")
	fmt.Println(r)
}
