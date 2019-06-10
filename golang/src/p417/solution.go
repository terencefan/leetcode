package main

type Grid int

func (g Grid) getX() int {
	return int(g&0xFF00) >> 8
}

func (g Grid) getY() int {
	return int(g & 0x00FF)
}

func New(x, y int) Grid {
	return Grid(x<<8 + y)
}

func build(height, width int) [][]int {
	r := make([][]int, height)
	for i := range r {
		r[i] = make([]int, width)
	}
	return r
}

var moves = [][]int{
	{1, 0},
	{-1, 0},
	{0, 1},
	{0, -1},
}

func bfs(matrix, continent [][]int) {
	height, width := len(continent), len(continent[0])

	q := make([]Grid, 0)
	for i := range continent {
		for j := range continent[i] {
			if continent[i][j] == 1 {
				q = append(q, New(i, j))
			}
		}
	}

	for len(q) > 0 {
		g := q[0]

		x, y := g.getX(), g.getY()
		for _, move := range moves {
			nx, ny := x+move[0], y+move[1]
			if nx == -1 || ny == -1 || nx == height || ny == width {
				continue
			}
			if matrix[nx][ny] < matrix[x][y] || continent[nx][ny] == 1 {
				continue
			}
			continent[nx][ny] = 1
			q = append(q, New(nx, ny))
		}

		q = q[1:]
	}
}

func pacificAtlantic(matrix [][]int) (r [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return [][]int{}
	}
	height, width := len(matrix), len(matrix[0])

	toPacific := build(height, width)
	toAtlantic := build(height, width)

	for i := 0; i < height; i++ {
		toPacific[i][0] = 1
		toAtlantic[i][width-1] = 1
	}

	for i := 0; i < width; i++ {
		toPacific[0][i] = 1
		toAtlantic[height-1][i] = 1
	}

	bfs(matrix, toPacific)
	bfs(matrix, toAtlantic)

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if toAtlantic[i][j] == 1 && toPacific[i][j] == 1 {
				r = append(r, []int{i, j})
			}
		}
	}
	return r
}
