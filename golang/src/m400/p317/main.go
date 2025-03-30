package p317

type Point struct {
	x, y int
}

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

const (
	EMPTY = iota
	BUILDING
	OBSTACLE
)

func bfs(grid, distance [][]int, i, j, house int) {
	q := []Point{{i, j}}

	step := 0
	for len(q) > 0 {
		step++
		length := len(q)

		for i := range length {
			point := q[i]

			for _, direction := range directions {
				nx, ny := point.x+direction[0], point.y+direction[1]
				if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) && grid[nx][ny] == -house {
					grid[nx][ny] = -house - 1
					distance[nx][ny] += step
					q = append(q, Point{nx, ny})
				}
			}
		}
		q = q[length:]
	}
}

func shortestDistance(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return -1
	}

	m, n := len(grid), len(grid[0])

	distance := make([][]int, m)
	for i := range m {
		distance[i] = make([]int, n)
	}

	houses := 0
	for i, row := range grid {
		for j, cell := range row {
			if cell == BUILDING {
				bfs(grid, distance, i, j, houses)
				houses++
			}
		}
	}

	r := 1 << 31
	for i, row := range grid {
		for j, cell := range row {
			if cell == -houses {
				r = min(r, distance[i][j])
			}
		}
	}
	if r == 1<<31 {
		return -1
	}
	return r
}
