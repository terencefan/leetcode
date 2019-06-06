package stdlib

type Stack struct {
	nodes  []interface{}
	length int
}

func NewStack(nodes ...interface{}) *Stack {
	return &Stack{
		nodes:  nodes,
		length: len(nodes),
	}
}

func (s *Stack) Push(node interface{}) {
	s.nodes = append(s.nodes[:s.length], node)
	s.length++
}

func (s *Stack) Pop() interface{} {
	defer func() {
		s.length--
	}()
	return s.Peek()
}

func (s *Stack) Peek() interface{} {
	if s.length == 0 {
		return nil
	}
	return s.nodes[s.length-1]
}

func (s *Stack) All() []interface{} {
	return s.nodes[:s.length]
}

func (s *Stack) Len() int {
	return s.length
}
