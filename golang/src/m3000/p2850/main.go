package p2850

type Grid [][]int

func copy(grid Grid) Grid {
	var res = make(Grid, len(grid))
	for i, row := range grid {
		res[i] = make([]int, len(row))
		for j, num := range row {
			res[i][j] = num
		}
	}
	return res
}

func hash(grid Grid) int {
	var res = 0
	for _, row := range grid {
		for _, num := range row {
			res = res*10 + num
		}
	}
	return res
}

var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func minimumMoves(grid Grid) int {
	q := make([]Grid, 0)
	q = append(q, grid)

	visited := make(map[int]bool)
	visited[hash(grid)] = true

	step := 0
	for len(q) > 0 {
		length := len(q)

		for k := range length {
			curr := q[k]
			modified := false
			for i, row := range curr {
				for j, val := range row {
					if val < 1 {
						continue
					}
					for _, d := range directions {
						nx, ny := i+d[0], j+d[1]
						if nx < 0 || nx >= 3 || ny < 0 || ny >= 3 || curr[nx][ny] >= curr[i][j] {
							continue
						}
						modified = true
						next := copy(curr)
						next[i][j]--
						next[nx][ny]++

						h := hash(next)
						if visited[h] {
							continue
						}
						visited[h] = true
						q = append(q, next)
					}
				}
			}

			if !modified {
				return step
			}
		}

		q = q[length:]
		step++
	}
	return -1
}
