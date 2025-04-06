package p934

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type Point struct {
	x, y int
}

func fill(grid [][]int, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != 1 {
		return
	}
	grid[i][j] = 2
	fill(grid, i-1, j)
	fill(grid, i+1, j)
	fill(grid, i, j-1)
	fill(grid, i, j+1)
}

func shortestBridge(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	for i := range m * n {
		if grid[i/n][i%n] == 1 {
			fill(grid, i/n, i%n)
			break
		}
	}

	q := make([]Point, 0)
	targets := make(map[Point]bool)

	for i, row := range grid {
		for j, val := range row {
			if val == 1 {
				grid[i][j] = -1
				q = append(q, Point{i, j})
			} else if val == 2 {
				grid[i][j] = 0
				targets[Point{i, j}] = true
			}
		}
	}

	miles := 0
	for len(q) > 0 {
		l := len(q)

		for i := range l {
			p := q[i]

			for _, d := range directions {
				nx, ny := p.x+d[0], p.y+d[1]
				if targets[Point{nx, ny}] {
					return miles
				}
				if nx < 0 || nx >= m || ny < 0 || ny >= n {
					continue
				} else if grid[nx][ny] == 0 {
					grid[nx][ny] = 1
					q = append(q, Point{nx, ny})
				} else if grid[nx][ny] > miles {
					grid[nx][ny] = miles
					q = append(q, Point{nx, ny})
				}
			}

		}

		q = q[l:]
		miles++
	}
	return -1
}
