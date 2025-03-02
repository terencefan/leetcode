package p994

const (
	EMPTY  = 0
	FRESH  = 1
	ROTTEN = 2
)

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func orangesRotting(grid [][]int) int {
	q := make([][]int, 0)

	for i, row := range grid {
		for j, val := range row {
			if val == ROTTEN {
				q = append(q, []int{i, j, 0})
			}
		}
	}

	var r = 0
	for len(q) > 0 {
		i, j, k := q[0][0], q[0][1], q[0][2]
		q = q[1:]

		if grid[i][j] == FRESH && k > r {
			r = k
		}
		grid[i][j] = ROTTEN

		for _, d := range directions {
			dx, dy := i+d[0], j+d[1]
			if dx >= 0 && dx < len(grid) && dy >= 0 && dy < len(grid[dx]) {
				if grid[dx][dy] == FRESH {
					q = append(q, []int{dx, dy, k + 1})
				}
			}
		}
	}

	for _, row := range grid {
		for _, val := range row {
			if val == FRESH {
				return -1
			}
		}
	}
	return r
}
