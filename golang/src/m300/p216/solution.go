package main

import (
	"fmt"
)

type Stack []int

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Pop() (r int) {
	r = (*s)[s.Len() - 1]
	*s = (*s)[:s.Len() - 1]
	return
}

var r [][]int

func dfs(s *Stack, k int, l int, remain int) {
	if 9 * k < remain {
		return
	}

	if k == 1 {
		s.Push(remain)
		t := make([]int, s.Len())
		for i, x := range *s {
			t[i] = x
		}
		r = append(r, t)
		s.Pop()
		return
	}

	for i := l; i * k < remain && i < 10; i++ {
		if i + 1 > remain - i {
			break
		}
		s.Push(i)
		dfs(s, k - 1, i + 1, remain - i)
		s.Pop()
	}
}

func combinationSum3(k int, n int) [][]int {
	r = make([][]int, 0) // just to make the codes easier...
	stack := make(Stack, 0, 10) // TODO maximum capacity can be optimized to sqrt n
	dfs(&stack, k, 1, n)
	return r
}

func main() {
	fmt.Println(combinationSum3(2, 18))
}
