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
)

var priorities = map[int]int{
	BGN: 0,
	ADD: 1,
	SUB: 1,
	MUL: 2,
	DIV: 2,
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
	return s[s.len()-1]
}

func (s *Stack) push(x int) {
	*s = append(*s, x)
}

func (s *Stack) pop() (r int) {
	r = s.peek()
	*s = (*s)[:s.len()-1]
	return
}

func calculate(s string) int {
	nums := make(Stack, 0)
	opts := make(Stack, 0)
	opts.push(BGN)

	num := 0
	for i := 0; i < len(s); i++ {
		c := int(s[i])
		if p, ok := priorities[c]; ok {
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

	for opts.peek() != BGN {
		num = functions[opts.pop()](nums.pop(), num)
	}
	return num
}

func main() {
	r := calculate("100000000/1/2/3/4/5/6/7/8/9/10")
	fmt.Println(r)
}
