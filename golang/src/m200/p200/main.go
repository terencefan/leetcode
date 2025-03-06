package p200

const ISLAND = '1'
const WATER = '0'

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func sink(grid [][]byte, x, y int) {
	if len(grid) == 0 {
		return
	} else if grid[x][y] == WATER {
		return
	}
	m, n := len(grid), len(grid[0])

	q := make([][]int, 0, m*n)
	q = append(q, []int{x, y})

	for len(q) > 0 {
		l := len(q)
		for i := range l {
			for _, d := range directions {
				x, y := q[i][0]+d[0], q[i][1]+d[1]
				if x >= 0 && x < m && y >= 0 && y < n && grid[x][y] == ISLAND {
					q = append(q, []int{x, y})
					grid[x][y] = WATER
				}
			}
		}
		q = q[l:]
	}
}

func numIslands(grid [][]byte) int {
	r := 0
	for i, row := range grid {
		for j, val := range row {
			if val == ISLAND {
				sink(grid, i, j)
				r++
			}
		}
	}
	return r
}
