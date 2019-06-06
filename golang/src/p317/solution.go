package main

import (
	"fmt"
)

const (
	EMPTY = iota
	BUILDING
	OBSTACLE
	VISITED
)

type Node int

func (n Node) getX() int {
	return int(n&0xFF00) >> 8
}

func (n Node) getY() int {
	return int(n & 0xFF)
}

func New(x, y int) Node {
	return Node(x<<8 + y)
}

type Solver struct {
	grid          [][]int
	height, width int
	buildings     []Node
}

func Construct(grid [][]int) Solver {
	return Solver{
		grid:      grid,
		height:    len(grid),
		width:     len(grid[0]),
		buildings: make([]Node, 0),
	}
}

func (s *Solver) findBuildings() (r int) {
	for i := range s.grid {
		for j := range s.grid[i] {
			if s.grid[i][j] == BUILDING {
				s.buildings = append(s.buildings, New(i, j))
				r++
			}
		}
	}
	return
}

func (s *Solver) isValid(x, y int) bool {
	return x >= 0 && y >= 0 && x < s.height && y < s.width
}

var ds = [][]int{
	{-1, 0},
	{1, 0},
	{0, -1},
	{0, 1},
}

func (s *Solver) bfs(start Node, cap int) (r int) {

	var copied = make([][]int, s.height)
	for i := range s.grid {
		copied[i] = make([]int, s.width)
		for j := range s.grid[i] {
			copied[i][j] = s.grid[i][j]
		}
	}

	q, steps, count := []Node{start}, 0, 0
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			node := q[i]
			x, y := node.getX(), node.getY()

			switch copied[x][y] {
			case OBSTACLE, VISITED:
				continue
			case BUILDING:
				count++
				r += steps
			case EMPTY:
				for _, d := range ds {
					nx, ny := x+d[0], y+d[1]
					if !s.isValid(nx, ny) {
						continue
					}
					q = append(q, New(nx, ny))
				}
			}
			copied[x][y] = VISITED
		}
		q = q[length:]
		steps++
	}

	if count != len(s.buildings) {
		return -1
	}
	return
}

func shortestDistance(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	s := Construct(grid)
	if s.findBuildings() == 0 {
		return -1
	}

	currentMin := 1<<32 - 1
	for i := range grid {
		for j := range grid[i] {
			if s.grid[i][j] != EMPTY {
				continue
			}
			r := s.bfs(New(i, j), currentMin)
			if r > 0 && r < currentMin {
				currentMin = r
			}
		}
	}
	if currentMin == 1<<32-1 {
		return -1
	}
	return currentMin
}

func main() {
	r := shortestDistance([][]int{
		{1, 1, 1, 1, 1, 0},
		{0, 0, 0, 0, 0, 1},
		{0, 1, 1, 0, 0, 1},
		{1, 0, 0, 1, 0, 1},
		{1, 0, 1, 0, 0, 1},
		{1, 0, 0, 0, 0, 1},
		{0, 1, 1, 1, 1, 0},
	})
	fmt.Println(r)
}
