package p227

type Stack[T int | byte | rune] struct {
	arr []T
}

func (s *Stack[T]) Push(x T) {
	s.arr = append(s.arr, x)
}

func (s *Stack[T]) Top() T {
	return s.arr[len(s.arr)-1]
}

func (s *Stack[T]) Pop() T {
	r := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return r
}

func (s *Stack[T]) Len() int {
	return len(s.arr)
}

func calc(op1, op2 int, op rune) int {
	switch op {
	case '+':
		return op1 + op2
	case '-':
		return op1 - op2
	case '*':
		return op1 * op2
	case '/':
		return op1 / op2
	}
	panic("invalid")
}

func priority(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	panic("invalid")
}

func calculate(s string) int {

	nums := Stack[int]{arr: make([]int, 0)}
	opts := Stack[rune]{arr: make([]rune, 0)}

	num := 0
	for _, c := range s {
		if c == ' ' {
			continue
		} else if c >= '0' && c <= '9' {
			num *= 10
			num += int(c - '0')
		} else {
			nums.Push(num)
			num = 0
			for opts.Len() > 0 {
				opt := opts.Pop()
				if priority(opt) >= priority(c) {
					op2, op1 := nums.Pop(), nums.Pop()
					nums.Push(calc(op1, op2, opt))
				} else {
					opts.Push(opt)
					break
				}
			}
			opts.Push(c)
		}
	}
	nums.Push(num)

	for opts.Len() > 0 {
		opt := opts.Pop()
		op2, op1 := nums.Pop(), nums.Pop()
		nums.Push(calc(op1, op2, opt))
	}
	return nums.Top()
}
