package main

import (
	"fmt"
)

type Stack []int

func (s *Stack) Empty() bool {
	return len(*s) == 0
}

func (s *Stack) Peek() int {
	return (*s)[len(*s)-1]
}

func (s *Stack) Pop() int {
	r := s.Peek()
	*s = (*s)[:len(*s)-1]
	return r
}

func (s *Stack) Push(x int) {
	*s = append(*s, x)
}

type MyQueue struct {
	s1, s2 *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	s1 := make(Stack, 0)
	s2 := make(Stack, 0)
	return MyQueue{&s1, &s2}
}

/** Push element x to the back of queue. */
func (q *MyQueue) Push(x int) {
	q.s1.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (q *MyQueue) Pop() int {
	r := q.Peek()
	q.s2.Pop()
	return r
}

/** Get the front element. */
func (q *MyQueue) Peek() int {
	if q.s2.Empty() {
		for !q.s1.Empty() {
			q.s2.Push(q.s1.Pop())
		}
	}
	return q.s2.Peek()
}

/** Returns whether the queue is empty. */
func (q *MyQueue) Empty() bool {
	return q.s1.Empty() && q.s2.Empty()
}

func main() {
	q := Constructor()
	q.Push(1)
	fmt.Println(q.Peek())
	q.Push(2)
	q.Push(3)
	fmt.Println(q.Peek())
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
