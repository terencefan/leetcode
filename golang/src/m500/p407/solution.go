package main

import (
	"container/heap"
	"fmt"
)

type Grid int

func (g Grid) getX() int {
	return int(g&0xFF000000) >> 24
}

func (g Grid) getY() int {
	return int(g&0x00FF0000) >> 16
}

func (g Grid) getHeight() int {
	return int(g & 0x0000FFFF)
}

func New(x, y, height int) Grid {
	return Grid(x<<24 + y<<16 + height)
}

type Heap []Grid

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	return h[i].getHeight() < h[j].getHeight()
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Heap) Push(x interface{}) {
	*h = append(*h, x.(Grid))
}

func (h *Heap) Pop() interface{} {
	r := (*h)[h.Len()-1]
	*h = (*h)[:h.Len()-1]
	return r
}

func (h Heap) Peek() Grid {
	return h[0]
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) < 3 || len(heightMap[0]) < 3 {
		return 0
	}
	height, width := len(heightMap), len(heightMap[0])

	var isValid = func(x, y int) bool {
		return x >= 0 && y >= 0 && x < height && y < width && heightMap[x][y] >= 0
	}

	var moves = [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	h := make(Heap, 0, 2*height+2*width-4)

	for i := 0; i < height; i++ {
		heap.Push(&h, New(i, 0, heightMap[i][0]))
		heap.Push(&h, New(i, width-1, heightMap[i][width-1]))
		heightMap[i][0] = -1
		heightMap[i][width-1] = -1
	}

	for i := 1; i < width-1; i++ {
		heap.Push(&h, New(0, i, heightMap[0][i]))
		heap.Push(&h, New(height-1, i, heightMap[height-1][i]))
		heightMap[0][i] = -1
		heightMap[height-1][i] = -1
	}

	currentMin, v := h.Peek().getHeight(), 0
	for h.Len() > 0 {
		node := heap.Pop(&h).(Grid)
		if node.getHeight() < currentMin {
			v += currentMin - node.getHeight()
		} else {
			currentMin = node.getHeight()
		}

		x, y := node.getX(), node.getY()
		for _, move := range moves {
			nx, ny := x+move[0], y+move[1]
			if isValid(nx, ny) {
				heap.Push(&h, New(nx, ny, heightMap[nx][ny]))
				heightMap[nx][ny] = -1
			}
		}
	}
	return v
}

func main() {
	r := trapRainWater([][]int{
		{1, 4, 3, 1, 3, 2},
		{3, 2, 1, 3, 2, 4},
		{2, 3, 3, 2, 3, 1},
	})
	fmt.Println(r)
}
