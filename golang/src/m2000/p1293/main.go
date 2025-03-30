package p1293

const (
	EMPTY = 0
	WALL  = 1
)

type State struct {
	x, y  int
	bombs int
}

var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func shortestPath(grid [][]int, k int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])

	q := []State{{0, 0, k}}

	var visited = make(map[State]bool)

	step := 0
	for len(q) > 0 {
		var length = len(q)

		for i := range length {
			s := q[i]

			if s.x == m-1 && s.y == n-1 {
				return step
			}

			for _, d := range directions {
				nx, ny := s.x+d[0], s.y+d[1]
				if nx >= 0 && nx < m && ny >= 0 && ny < n {
					if grid[nx][ny] == WALL && s.bombs > 0 {
						next := State{nx, ny, s.bombs - 1}
						if !visited[next] {
							q = append(q, next)
							visited[next] = true
						}
					} else if grid[nx][ny] == EMPTY {
						next := State{nx, ny, s.bombs}
						if !visited[next] {
							q = append(q, next)
							visited[next] = true
						}
					}
				}
			}
		}
		q = q[length:]
		step++
	}
	return -1
}
