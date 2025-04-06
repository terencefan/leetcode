package p856

type IndexRune struct {
	idx int
	b   rune
}

type IndexNum struct {
	idx, num int
}

type Stack[T any] struct {
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

func scoreOfParentheses(s string) int {

	var ps = Stack[IndexRune]{arr: make([]IndexRune, 0)}
	var ns = Stack[IndexNum]{arr: make([]IndexNum, 0)}

	for i, c := range s {
		switch c {
		case '(':
			ps.Push(IndexRune{i, c})
		case ')':
			p := ps.Pop()
			v := 0

			for ns.Len() > 0 {
				n := ns.Pop()
				if n.idx < p.idx {
					ns.Push(n)
					break
				}
				v += n.num
			}

			if v == 0 {
				ns.Push(IndexNum{p.idx, 1})
			} else {
				ns.Push(IndexNum{p.idx, 2 * v})
			}
		}
	}

	r := 0
	for _, v := range ns.arr {
		r += v.num
	}
	return r
}
