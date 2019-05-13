package stdlib

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
