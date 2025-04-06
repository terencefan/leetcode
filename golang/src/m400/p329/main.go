package p329

var directions = [][]int{{1, 0}, {-1, 0}, {0, -1}, {0, 1}}

type Point struct {
	x, y, step int
}

func dfs(matrix, cache [][]int, x, y int) {
	cache[x][y] = 1

	s := []Point{{x, y, 1}}

	var p Point
	for len(s) > 0 {
		p, s = s[len(s)-1], s[:len(s)-1]

		for _, d := range directions {
			nx, ny := p.x+d[0], p.y+d[1]
			if nx < 0 || nx >= len(matrix) || ny < 0 || ny >= len(matrix[0]) {
				continue
			}
			if matrix[nx][ny] <= matrix[p.x][p.y] {
				continue
			}
			if p.step+1 > cache[nx][ny] {
				cache[nx][ny] = p.step + 1
				s = append(s, Point{nx, ny, p.step + 1})
			}
		}
	}
}

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	m, n := len(matrix), len(matrix[0])

	cache := make([][]int, m)
	for i := range m {
		cache[i] = make([]int, n)
	}

	for i := range m {
		for j := range n {
			if cache[i][j] != 0 {
				continue
			}
			dfs(matrix, cache, i, j)
		}
	}

	r := 0
	for i := range m {
		for j := range n {
			r = max(r, cache[i][j])
		}
	}
	return r
}
