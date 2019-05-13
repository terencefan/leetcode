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

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func MakeTreeNode(v int) (node *TreeNode) {
	node = &TreeNode{}
	node.Val = v
	return
}

func walk(root *TreeNode) (t []int) {
	var s1 = NewStack(root)
	var s2 = NewStack()

	for {
		if s1.Len() > 0 {
			var node = s1.Pop().(*TreeNode)
			if node.Left != nil {
				s1.Push(node.Left)
			}
			s2.Push(node)
		} else if s2.Len() > 0 {
			var node = s2.Pop().(*TreeNode)
			if node.Right != nil {
				s1.Push(node.Right)
			}
			t = append(t, node.Val)
		} else {
			break
		}
	}
	return t
}

func GetMinimumDifference(root *TreeNode) (r int) {
	var t = walk(root)

	r = t[1] - t[0]
	for i := 2; i < len(t); i++ {
		if t[i]-t[i-1] < r {
			r = t[i] - t[i-1]
		}
	}
	return
}
