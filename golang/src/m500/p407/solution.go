package p407

import "container/heap"

type Block struct {
	x, y, h int
}

type BlockHeap []Block

func (h BlockHeap) Len() int           { return len(h) }
func (h BlockHeap) Less(i, j int) bool { return h[i].h < h[j].h }
func (h BlockHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *BlockHeap) Push(x any) {
	*h = append(*h, x.(Block))
}

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (h *BlockHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func trapRainWater(heightMap [][]int) int {
	if len(heightMap) == 0 {
		return 0
	}
	m, n := len(heightMap), len(heightMap[0])

	bounder := &BlockHeap{}
	heap.Init(bounder)

	for i := 0; i < m; i++ {
		heap.Push(bounder, Block{i, 0, heightMap[i][0]})
		heightMap[i][0] = -1
		heap.Push(bounder, Block{i, n - 1, heightMap[i][n-1]})
		heightMap[i][n-1] = -1
	}
	for j := 1; j < n; j++ {
		heap.Push(bounder, Block{0, j, heightMap[0][j]})
		heightMap[0][j] = -1
		heap.Push(bounder, Block{m - 1, j, heightMap[m-1][j]})
		heightMap[m-1][j] = -1
	}

	var r = 0
	for bounder.Len() > 0 {
		b := heap.Pop(bounder).(Block)

		for _, d := range directions {
			x, y := b.x+d[0], b.y+d[1]
			if x >= 0 && x < m && y >= 0 && y < n && heightMap[x][y] != -1 {
				if b.h > heightMap[x][y] {
					r += b.h - heightMap[x][y]
				}
				heap.Push(bounder, Block{x, y, max(b.h, heightMap[x][y])})
				heightMap[x][y] = -1
			}
		}
	}
	return r
}
