package main

import (
	"fmt"
)

func sort(buildings [][]int, l, r int) {
	if l >= r {
		return
	}

	p, pivot := l, buildings[l]
	i, j := l, r
	for i < j {
		for ;i<j;j-- {
			if buildings[j][0] < pivot[0] {
				buildings[p] = buildings[j]
				p = j
				break
			}
		}
		for ;i<j;i++ {
			if buildings[i][0] > pivot[0] {
				buildings[p] = buildings[i]
				p = i
				break
			}
		}
	}
	buildings[i] = pivot

	sort(buildings, l, i - 1)
	sort(buildings, i + 1, r)
}

type Heap [][]int

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return (h)[i][2] > (h)[j][2]
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h Heap) Peek() []int {
	return h[0]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *Heap) Pop() interface{} {
	r := (*h)[h.Len() - 1]
	*h = (*h)[:h.Len() - 1]
	return r
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

type skyline [][]int

func (s *skyline) push(x, y int) {
	lp := len(*s)
	if lp > 0 {
		if x == (*s)[lp - 1][0] {
			(*s)[lp - 1][1] = y
		} else if y == (*s)[lp - 1][1] {
			return
		} else {
			*s = append(*s, []int{x, y})
		}
	} else {
		*s = append(*s, []int{x, y})
	}
}

func mergeSkyline(l, r [][]int) (p skyline) {
	i, j := 0, 0

	x, ly, ry := 0, 0, 0

	for i < len(l) && j < len(r) {
		if l[i][0] < r[j][0] {
			x = l[i][0]
			ly = l[i][1]
			i++
		} else {
			x = r[j][0]
			ry = r[j][1]
			j++
		}

		y := max(ly, ry)

		p.push(x, y)

	}

	for _, t := range l[i:] {
		p.push(t[0], t[1])
	}
	for _, t := range r[j:] {
		p.push(t[0], t[1])
	}
	return p
}

func getSkyline(buildings [][]int) (r [][]int) {
	if len(buildings) == 0 {
		return [][]int{}
	} else if len(buildings) == 1 {
		return [][]int{
			{buildings[0][0], buildings[0][2]},
			{buildings[0][1], 0},
		}
	}

	m := len(buildings) / 2

	left := getSkyline(buildings[:m])
	right := getSkyline(buildings[m:])
	return mergeSkyline(left, right)
}

func main() {
	// r := getSkyline([][]int{
	// 	{5, 12, 12},
	// 	{3, 7, 15},
	// 	{19, 24, 8},
	// 	{15, 20, 10},
	// 	{2, 9, 10},
	// })
	r := getSkyline([][]int{
		{1,2,1},
		{1,2,2},
		{1,2,3},
	})
	fmt.Println(r)
}
