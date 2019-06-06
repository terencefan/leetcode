package main

import (
	"container/heap"
	"fmt"
)

func quicksort(arr []int, l, r int) {
	if l >= r {
		return
	}

	p, pivot := l, arr[l]

	i, j := l, r
	for i < j {
		for ; i < j; j-- {
			if arr[j] < pivot {
				arr[p] = arr[j]
				p = j
				break
			}
		}
		for ; i < j; i++ {
			if arr[i] > pivot {
				arr[p] = arr[i]
				p = i
				break
			}
		}
	}
	arr[p] = pivot

	quicksort(arr, l, p-1)
	quicksort(arr, p+1, r)
}

type Heap [][]int

func (h *Heap) Len() int {
	return len(*h)
}

func (h *Heap) Less(i, j int) bool {
	b1, b2 := (*h)[i], (*h)[j]
	return b1[2] > b2[2]
}

func (h *Heap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *Heap) Peek() []int {
	return (*h)[0]
}

func (h *Heap) Pop() interface{} {
	l := h.Len()
	r := (*h)[l-1]
	*h = (*h)[:l-1]
	return r
}

func getSkyline(buildings [][]int) [][]int {

	lines := make([]int, 0)
	for _, building := range buildings {
		lines = append(lines, building[0])
		lines = append(lines, building[1])
	}
	quicksort(lines, 0, len(lines)-1)

	buildings = append(buildings, []int{1<<32 - 1, 1<<32 - 1})

	h := make(Heap, 0)
	h.Push([]int{0, 1<<32 - 1, 0})

	i, skyline := 0, [][]int{{0, 0}}
	for _, line := range lines {
		for i < len(buildings) && line == buildings[i][0] {
			heap.Push(&h, buildings[i])
			i++
		}

		for h.Len() > 0 && line >= h.Peek()[1] {
			heap.Pop(&h)
		}

		if h.Peek()[2] == skyline[len(skyline)-1][1] {
			continue
		}
		skyline = append(skyline, []int{line, h.Peek()[2]})
	}
	return skyline[1:]
}

func main() {
	r := getSkyline([][]int{
		{2, 9, 10},
		{3, 7, 15},
		{5, 12, 12},
		{15, 20, 10},
		{19, 24, 8},
	})
	fmt.Println(r)
	// r := getSkyline([][]int{
	// 	{1,2,1},
	// 	{1,2,2},
	// 	{1,2,3},
	// })
	// fmt.Println(r)
}
