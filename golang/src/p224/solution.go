package main

import (
	"fmt"
)

type Stack []int

func (s *Stack) push(x int) {
	*s = append(*s, x)
}

func (s *Stack) peek() (r int) {
	return (*s)[s.len()-1]
}

func (s *Stack) pop() (r int) {
	r = s.peek()
	*s = (*s)[:s.len()-1]
	return r
}

func (s *Stack) len() int {
	return len(*s)
}

func calc(a, b, opt int) int {
	switch opt {
	case '-':
		return a - b
	case '+':
		return a + b
	}
	return 0
}

func calculate(s string) (r int) {

	nums := new(Stack)
	opts := new(Stack)

	num := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			num *= 10
			num += int(c - '0')
		case '-', '+':
			for opts.len() > 0 && opts.peek() != '(' {
				num = calc(nums.pop(), num, opts.pop())
			}
			nums.push(num)
			opts.push(int(c))
			num = 0
		case ')':
			for opts.len() > 0 && opts.peek() != '(' {
				num = calc(nums.pop(), num, opts.pop())
			}
			opts.pop()
		case '(':
			opts.push(int(c))
		}
	}
	for opts.len() > 0 {
		num = calc(nums.pop(), num, opts.pop())
	}
	return num
}

func main() {
	r := calculate("2 - 1 + 2")
	fmt.Println(r)
}
