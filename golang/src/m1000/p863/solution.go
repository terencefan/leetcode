package main

import (
	"fmt"
)

type Solver struct {
	stack []*TreeNode
}

func (s *Solver) find(node, target *TreeNode) bool {
	if node == nil {
		return false
	} else if node == target {
		return true
	}

	s.stack = append(s.stack, node)
	if s.find(node.Left, target) {
		return true
	}
	if s.find(node.Right, target) {
		return true
	}
	s.stack = s.stack[:len(s.stack)-1]
	return false
}

func (s *Solver) bfs(node *TreeNode, K int) []*TreeNode {
	q := []*TreeNode{node}
	for K > 0 && len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		q = q[length:]
		K--
	}
	return q
}

func distanceK(root *TreeNode, target *TreeNode, K int) []int {
	if target == nil {
		return []int{}
	} else if K == 0 {
		return []int{target.Val}
	}

	s := &Solver{}
	if !s.find(root, target) {
		return []int{}
	}

	var nodes = s.bfs(target, K)

	next := target
	for len(s.stack) > 0 && K >= 2 {
		parent := s.stack[len(s.stack)-1]
		s.stack = s.stack[:len(s.stack)-1]

		if parent.Left != nil && parent.Left != next {
			nodes = append(nodes, s.bfs(parent.Left, K-2)...)
		}
		if parent.Right != nil && parent.Right != next {
			nodes = append(nodes, s.bfs(parent.Right, K-2)...)
		}
		next = parent
		K--
	}

	fmt.Println(s.stack)
	if len(s.stack) >= K {
		nodes = append(nodes, s.stack[len(s.stack)-K])
	}

	var r []int
	for _, node := range nodes {
		r = append(r, node.Val)
	}
	return r
}
