package p772

const (
	MUL = '*'
	DIV = '/'
	ADD = '+'
	SUB = '-'
	LPR = '('
	RPR = ')'
)

func priority(opt rune) int {
	switch opt {
	case ADD, SUB:
		return 1
	case MUL, DIV:
		return 2
	case LPR, RPR:
		return -1
	}
	panic("invalid")
}

func calc(opt rune, op1, op2 int) int {
	switch opt {
	case ADD:
		return op1 + op2
	case SUB:
		return op1 - op2
	case MUL:
		return op1 * op2
	case DIV:
		return op1 / op2
	}
	panic("invalid")
}

type Stack[T int | rune] struct {
	arr []T
}

func (s *Stack[T]) Push(x T) {
	s.arr = append(s.arr, x)
}

func (s *Stack[T]) Pop() T {
	r := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return r
}

func (s *Stack[T]) Len() int {
	return len(s.arr)
}

func calculate(s string) int {

	opts := Stack[rune]{arr: make([]rune, 0)}
	nums := Stack[int]{arr: make([]int, 0)}

	num := 0
	for _, c := range s {
		if c == LPR {
			opts.Push(c)
		} else if c == RPR {
			nums.Push(num)
			num = 0
			for opts.Len() > 0 {
				opt := opts.Pop()
				if opt == LPR {
					num = nums.Pop()
					break
				}
				op2, op1 := nums.Pop(), nums.Pop()
				nums.Push(calc(opt, op1, op2))
			}
		} else if c >= '0' && c <= '9' {
			num *= 10
			num += int(c - '0')
		} else {
			nums.Push(num)
			num = 0
			for opts.Len() > 0 {
				opt := opts.Pop()
				if priority(c) <= priority(opt) {
					op2, op1 := nums.Pop(), nums.Pop()
					nums.Push(calc(opt, op1, op2))
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
		nums.Push(calc(opt, op1, op2))
	}
	return nums.Pop()
}
