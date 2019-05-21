package main

import (
	"container/heap"
	"fmt"
)

func compare(a, b []int) int {
	if a[0] < b[0] {
		return -1
	} else if a[0] == b[0] {
		if a[1] < b[1] {
			return -1
		} else if a[1] == b[1] {
			return 0
		} else {
			return 1
		}
	} else {
		return 1
	}
}

func sort(intervals [][]int, l, r int) {
	if l >= r {
		return
	}

	pivot := intervals[l]
	p := l

	i, j := l, r
	for i < j {
		for ;j>i;j-- {
			if compare(intervals[j], pivot) < 0 {
				intervals[p] = intervals[j]
				p = j
				break
			}
		}
		for ;i<j;i++ {
			if compare(intervals[i], pivot) > 0 {
				intervals[p] = intervals[i]
				p = i
				break
			}
		}
	}
	intervals[p] = pivot

	sort(intervals, l, p - 1)
	sort(intervals, p + 1, r)
}

type endtime struct {
	arr []int
}

func (e *endtime) Less(i, j int) bool {
	return e.arr[i] < e.arr[j]
}

func (e *endtime) Push(x interface{}) {
	e.arr = append(e.arr, x.(int))
}

func (e *endtime) Swap(i, j int) {
	e.arr[i], e.arr[j] = e.arr[j], e.arr[i]
}

func (e *endtime) Peek() int {
	return e.arr[0]
}

func (e *endtime) Pop() (r interface{}) {
	r = e.arr[len(e.arr) - 1]
	e.arr = e.arr[:len(e.arr) - 1]
	return
}

func (e *endtime) Len() int {
	return len(e.arr)
}

func minMeetingRooms(intervals [][]int) int {
	sort(intervals, 0, len(intervals) - 1)

	h, r := &endtime{}, 0
	for _, interval := range intervals {
		for h.Len() > 0 && h.Peek() <= interval[0] {
			heap.Pop(h)
		}
		heap.Push(h, interval[1])
		if h.Len() > r {
			r = h.Len()
		}
	}
	return r
}

func main() {
	r := minMeetingRooms([][]int{
		{13, 15},
		{1, 13},
	})
	fmt.Println(r)
}
