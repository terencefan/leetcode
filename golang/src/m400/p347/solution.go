package main

import (
	"container/heap"
	"fmt"
)

type Node struct {
	num, count int
}

type IntArray struct {
	arr []*Node
}

func (a *IntArray) Len() int {
	return len(a.arr)
}

func (a *IntArray) Less(i, j int) bool {
	return a.arr[i].count < a.arr[j].count
}

func (a *IntArray) Swap(i, j int) {
	a.arr[i], a.arr[j] = a.arr[j], a.arr[i]
}

func (a *IntArray) Push(i interface{}) {
	if x, ok := i.(*Node); ok {
		a.arr = append(a.arr, x)
	}
}

func (a *IntArray) Pop() (r interface{}) {
	r = a.arr[a.Len()-1]
	a.arr = a.arr[:a.Len()-1]
	return
}

func topKFrequent(nums []int, k int) (r []int) {
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	arr := &IntArray{make([]*Node, 0)}
	heap.Init(arr)

	for num, count := range m {
		heap.Push(arr, &Node{num, count})
		if arr.Len() > k {
			heap.Pop(arr)
		}
	}

	for _, node := range arr.arr {
		r = append(r, node.num)
	}
	return
}

func main() {
	r := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Print(r)
}
