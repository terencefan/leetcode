package main

import (
	"fmt"
)

type graph struct {
	matrix [][]int
	length [][]int
	max    int
}

func makeGraph(matrix [][]int) graph {
	var length = make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		length[i] = make([]int, len(matrix[0]))
	}
	return graph{
		matrix: matrix,
		length: length,
	}
}

func (g graph) height() int {
	return len(g.matrix)
}

func (g graph) width() int {
	return len(g.matrix[0])
}

func (g *graph) dfs(x, y, l int) {
	if l <= g.length[x][y] {
		return
	}
	g.length[x][y] = l
	if l > g.max {
		g.max = l
	}

	val := g.matrix[x][y]
	if x > 0 && g.matrix[x-1][y] > val {
		g.dfs(x-1, y, l+1)
	}
	if y > 0 && g.matrix[x][y-1] > val {
		g.dfs(x, y-1, l+1)
	}
	if x < g.height()-1 && g.matrix[x+1][y] > val {
		g.dfs(x+1, y, l+1)
	}
	if y < g.width()-1 && g.matrix[x][y+1] > val {
		g.dfs(x, y+1, l+1)
	}
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	g := makeGraph(matrix)
	for i := 0; i < g.height(); i++ {
		for j := 0; j < g.width(); j++ {
			g.dfs(i, j, 1)
		}
	}
	return g.max
}

func main() {
	r := longestIncreasingPath([][]int{
		{9, 9, 4},
		{6, 6, 8},
		{2, 1, 1},
	})
	fmt.Println(r)
}
