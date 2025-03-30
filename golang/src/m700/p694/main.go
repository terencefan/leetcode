package p694

var directions = [][]int{{0, 1, 'd'}, {0, -1, 'u'}, {1, 0, 'r'}, {-1, 0, 'l'}}

func hash(grid [][]int, i, j int) string {
	q := [][]int{{i, j}}
	bytes := []byte{}
	step := 0
	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		for _, d := range directions {
			nx, ny := x+d[0], y+d[1]
			bytes = append(bytes, byte(step))

			if nx >= 0 && nx < len(grid) && ny >= 0 && ny < len(grid[0]) && grid[nx][ny] == 1 {
				grid[nx][ny] = 0
				bytes = append(bytes, byte(d[2]))
				q = append(q, []int{nx, ny})
			}
		}
		step++
		q = q[1:]
	}
	return string(bytes)
}

func numDistinctIslands(grid [][]int) int {
	var r = make(map[string]bool)
	for i, row := range grid {
		for j, v := range row {
			if v == 1 {
				var h = hash(grid, i, j)
				r[h] = true
			}
		}
	}
	return len(r)
}
