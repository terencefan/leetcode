package main

type Stack struct {
	nodes []interface{}
	len   int
}

func stack(nodes ...interface{}) *Stack {
	var s = Stack{}
	s.nodes = nodes
	s.len = len(nodes)
	return &s
}

func (s *Stack) push(node interface{}) {
	s.nodes = append(s.nodes[:s.len], node)
	s.len++
}

func (s *Stack) pop() interface{} {
	if s.len == 0 {
		return nil
	}
	s.len--
	return s.nodes[s.len]
}

func (s *Stack) back() interface{} {
	if s.len == 0 {
		return nil
	}
	return s.nodes[s.len-1]
}

func (s *Stack) all() []interface{} {
	return s.nodes[:s.len]
}
