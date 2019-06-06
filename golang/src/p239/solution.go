package main

import (
	"fmt"
)

type Deque struct {
	l, r int
	arr  []int
	size int
}

func mod(x, mod int) int {
	r := x % mod
	if r < 0 {
		return r + mod
	} else {
		return r
	}
}

func (q *Deque) pushBack(x int) {
	q.arr[mod(q.r, q.size)] = x
	q.r++
}

func (q *Deque) pushFront(x int) {
	q.arr[mod(q.l, q.size)] = x
	q.l--
}

func (q *Deque) back() int {
	return q.arr[mod(q.r-1, q.size)]
}

func (q *Deque) front() int {
	return q.arr[mod(q.l+1, q.size)]
}

func (q *Deque) popBack() {
	q.r--
}

func (q *Deque) popFront() {
	q.l++
}

func (q *Deque) len() int {
	return q.r - q.l - 1
}

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) < k || k <= 0 {
		return []int{}
	}

	r := make([]int, 0)

	q := &Deque{0, 1, make([]int, k), k}

	for i := 0; i < k; i++ {
		for q.len() > 0 && nums[i] > nums[q.back()] {
			q.popBack()
		}
		q.pushBack(i)
	}

	for i := 0; i < len(nums)-k; i++ {
		// fmt.Println(i, q.front(), nums[q.front()])
		r = append(r, nums[q.front()])

		for q.len() > 0 && nums[i+k] > nums[q.back()] {
			// fmt.Println("back", q.back(), nums[q.back()])
			q.popBack()
		}

		for q.len() > 0 && q.front() <= i {
			// fmt.Println("front", q.front(), nums[q.front()])
			q.popFront()
		}
		q.pushBack(i + k)
	}
	r = append(r, nums[q.front()])

	return r
}

func main() {
	r := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	fmt.Println(r)
}
