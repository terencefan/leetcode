package main

import (
	"fmt"
)

type Stack []int

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Peek() int {
	return (*s)[s.Len()-1]
}

func (s *Stack) Pop() int {
	r := s.Peek()
	*s = (*s)[:s.Len()-1]
	return r
}

const INTMIN = -(1 << 32)

func nextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	var s = make(Stack, 0)
	var r = make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		r[i] = INTMIN
	}

	for i := len(nums) - 1; i >= 0; i-- {
		for s.Len() > 0 && nums[i] >= nums[s.Peek()] {
			s.Pop()
		}

		if s.Len() > 0 {
			r[i] = nums[s.Peek()]
		}
		s.Push(i)
	}

	for i := len(nums) - 1; i >= 0; i-- {
		if r[i] != INTMIN {
			continue
		}
		for s.Len() > 0 && nums[i] >= nums[s.Peek()] {
			s.Pop()
		}

		if s.Len() > 0 {
			r[i] = nums[s.Peek()]
		}
		s.Push(i)
	}

	for i := 0; i < len(nums); i++ {
		if r[i] == INTMIN {
			r[i] = -1
		}
	}
	return r
}

func main() {
	fmt.Println(nextGreaterElements([]int{-1, 0}))
}
