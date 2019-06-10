package main

import (
	"fmt"
)

const (
	EMPTY   = 0
	WALL    = 1
	VISITED = 2
)

type Node int

func (n Node) getX() int {
	return int(n&0xFF00) >> 8
}

func (n Node) getY() int {
	return int(n & 0x00FF)
}

func New(x, y int) Node {
	return Node(x<<8 + y)
}

var moves = [][]int{
	{-1, 0},
	{1, 0},
	{0, 1},
	{0, -1},
}

func hasPath(maze [][]int, start []int, destination []int) bool {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return false
	}
	height, width := len(maze), len(maze[0])

	var isWall = func(x, y int) bool {
		return x == -1 || y == -1 || x == height || y == width || maze[x][y] == WALL
	}

	if isWall(start[0], start[1]) || isWall(destination[0], destination[1]) {
		return false
	}
	var node, dest = New(0, 0), New(destination[0], destination[1])
	var q = []Node{New(start[0], start[1])}

	for len(q) > 0 {
		node, q = q[0], q[1:]
		if node == dest {
			return true
		}

		x, y := node.getX(), node.getY()
		if maze[x][y] == VISITED {
			continue
		}
		maze[x][y] = VISITED

		for _, move := range moves {
			k, nx, ny := 0, x+move[0], y+move[1]
			for k = 0; !isWall(nx, ny); k++ {
				nx += move[0]
				ny += move[1]
			}
			if k > 0 {
				nx -= move[0]
				ny -= move[1]
				q = append(q, New(nx, ny))
			}
		}
	}

	return false
}

func main() {
	r := hasPath(
		[][]int{
			{0, 0, 1, 0, 0},
			{0, 0, 0, 0, 0},
			{0, 0, 0, 1, 0},
			{1, 1, 0, 1, 1},
			{0, 0, 0, 0, 0},
		},
		[]int{0, 4},
		[]int{4, 4},
	)
	fmt.Println(r)
}
