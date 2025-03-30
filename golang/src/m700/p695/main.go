package p695

const (
	SEA  = 0
	LAND = 1
)

var direction = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func sink(grid [][]int, i, j int) int {
	if grid[i][j] == SEA {
		return 0
	}
	grid[i][j] = SEA

	size, q := 0, [][]int{{i, j}}

	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		for _, d := range direction {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) && grid[nx][ny] == LAND {
				q = append(q, []int{nx, ny})
				grid[nx][ny] = SEA
			}
		}
		size++
		q = q[1:]
	}
	return size
}

func maxAreaOfIsland(grid [][]int) int {
	r := 0
	for i, row := range grid {
		for j, _ := range row {
			r = max(r, sink(grid, i, j))
		}
	}
	return r
}
