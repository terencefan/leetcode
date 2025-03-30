package p934

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

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

func sail(grid [][]int, i, j int) {
	q := [][]int{{i, j}}

	miles := 0
	for len(q) > 0 {
		l := len(q)
		miles--

		for _, p := range q {
			x, y := p[0], p[1]
			for _, d := range directions {
				nx, ny := x+d[0], y+d[1]
				if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) {
					if grid[nx][ny] > 0 {
						continue
					} else if grid[nx][ny] == 0 || miles > grid[nx][ny] {
						grid[nx][ny] = miles
						q = append(q, []int{nx, ny})
					}
				}
			}
		}
		q = q[l:]
	}
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

	for i, row := range grid {
		for j, val := range row {
			if val == 2 {
				sail(grid, i, j)
			}
		}
	}

	return 0
}
