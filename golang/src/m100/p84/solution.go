package main

import (
	"fmt"
)

type Node struct {
	pos, height int
}

type Stack []Node

func (s *Stack) Push(x Node) {
	*s = append(*s, x)
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Peek() Node {
	return (*s)[s.Len()-1]
}

func (s *Stack) Pop() (r Node) {
	r = s.Peek()
	*s = (*s)[:s.Len()-1]
	return
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func largestRectangleArea(heights []int) (r int) {
	heights = append(heights, 0)

	s := make(Stack, 0)

	for i, height := range heights {
		for s.Len() > 0 && s.Peek().height > height {
			node := s.Pop()
			if s.Len() > 0 {
				r = max(r, node.height*(i-s.Peek().pos-1))
				fmt.Println(i, r, s.Peek().pos)
			} else {
				r = max(r, node.height*i)
			}
		}
		s.Push(Node{i, height})
	}
	return
}

func main() {
	r := largestRectangleArea([]int{4, 2, 0, 3, 2, 5})
	fmt.Println(r)
}
