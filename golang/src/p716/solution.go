package main

import "fmt"

type MaxStack struct {
	arr []int
	max []int
}

/** initialize your data structure here. */
func Constructor() MaxStack {
	return MaxStack{
		arr: make([]int, 0),
		max: make([]int, 0),
	}
}

func (s *MaxStack) Len() int {
	return len(s.arr)
}

func (s *MaxStack) Push(x int) {
	if s.Len() == 0 {
		s.max = append(s.max, x)
	} else {
		pMax := s.PeekMax()
		if x > pMax {
			s.max = append(s.max, x)
		} else {
			s.max = append(s.max, pMax)
		}
	}
	s.arr = append(s.arr, x)
}

func (s *MaxStack) Pop() (r int) {
	r = s.Top()
	s.arr = s.arr[:len(s.arr)-1]
	s.max = s.max[:len(s.max)-1]
	return
}

func (s *MaxStack) Top() (r int) {
	if s.Len() == 0 {
		return -1
	}
	r = s.arr[s.Len()-1]
	return
}

func (s *MaxStack) PeekMax() (r int) {
	if s.Len() == 0 {
		return -1
	}
	r = s.max[s.Len()-1]
	return
}

func (s *MaxStack) PopMax() int {
	pMax := s.PeekMax()

	buf := make([]int, 0)
	for {
		t := s.Pop()
		if pMax == t {
			break
		} else {
			buf = append(buf, t)
		}
	}

	for i := len(buf) - 1; i >= 0; i-- {
		s.Push(buf[i])
	}

	return pMax
}

func main() {
	s := Constructor()
	s.Push(5)
	s.Push(1)
	s.Push(5)
	fmt.Println(s.Top())
}
