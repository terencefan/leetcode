package main

const (
	GATE = 0
	WALL = -1
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
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func bfs(rooms [][]int, start Node) {
	height, width := len(rooms), len(rooms[0])

	q := []Node{start}

	steps := 0
	for len(q) > 0 {
		length := len(q)
		for i := 0; i < length; i++ {
			node := q[i]
			x, y := node.getX(), node.getY()
			if rooms[x][y] <= steps && steps != 0 {
				continue
			}
			rooms[x][y] = steps

			for _, move := range moves {
				nx, ny := x+move[0], y+move[1]
				if nx < 0 || ny < 0 || nx == height || ny == width {
					continue
				} else if rooms[nx][ny] == WALL {
					continue
				}
				q = append(q, New(nx, ny))
			}
		}
		q = q[length:]
		steps++
	}
}

func wallsAndGates(rooms [][]int) {
	if len(rooms) == 0 || len(rooms[0]) == 0 {
		return
	}

	for i := 0; i < len(rooms); i++ {
		for j := 0; j < len(rooms[0]); j++ {
			if rooms[i][j] == GATE {
				bfs(rooms, New(i, j))
			}
		}
	}
}
