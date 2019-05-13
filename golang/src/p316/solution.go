package main

type Stack struct {
	nodes []interface{}
	len   int
}

func NewStack(nodes ...interface{}) *Stack {
	var s = Stack{}
	s.nodes = nodes
	s.len = len(nodes)
	return &s
}

func (s *Stack) Push(node interface{}) {
	s.nodes = append(s.nodes[:s.len], node)
	s.len++
}

func (s *Stack) Pop() interface{} {
	if s.len == 0 {
		return nil
	}
	s.len--
	return s.nodes[s.len]
}

func (s *Stack) Back() interface{} {
	if s.len == 0 {
		return nil
	}
	return s.nodes[s.len-1]
}

func (s *Stack) All() []interface{} {
	return s.nodes[:s.len]
}

func (s *Stack) Len() int {
	return s.len
}
func removeDuplicateLetters(s string) string {
	var a = [256]int{}
	var b = [256]bool{}

	for _, x := range s {
		a[x]++
	}

	var st = NewStack(byte('_'))

	for _, x := range s {
		a[x]--
		if b[x] {
			continue
		}
		var back = st.Back().(byte)
		for byte(x) < back && a[back] > 0 {
			b[back] = false
			st.Pop()
			back = st.Back().(byte)
		}
		st.Push(byte(x))
		b[x] = true
	}

	var res = make([]byte, 32)
	for _, c := range st.All()[1:] {
		res = append(res, c.(byte))
	}
	return string(res)
}

func RemoveDuplicateLetters(s string) string {
	return removeDuplicateLetters(s)
}
