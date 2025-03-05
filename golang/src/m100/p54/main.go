package p54

var directions = [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}

const INTMAX = 1 << 31

func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	m, n := len(matrix), len(matrix[0])

	var r = make([]int, m*n)
	r[0] = matrix[0][0]
	matrix[0][0] = INTMAX

	k, d := 1, 0
	i, j := 0, 0
	for k < m*n {
		direction := directions[d%4]
		dx, dy := direction[0], direction[1]
		i, j = i+dx, j+dy

		for ; i >= 0 && i < m && j >= 0 && j < n && matrix[i][j] != INTMAX; i, j = i+dx, j+dy {
			r[k] = matrix[i][j]
			matrix[i][j] = INTMAX
			k++
		}
		i -= dx
		j -= dy
		d += 1
	}
	return r
}
