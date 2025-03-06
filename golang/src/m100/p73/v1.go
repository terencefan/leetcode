package p73

func setZeroes(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])

	var row0, col0 bool

	for i := range m {
		if matrix[i][0] == 0 {
			col0 = true
			break
		}
	}

	for i := range n {
		if matrix[0][i] == 0 {
			row0 = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := range m {
		for j := range n {
			x, y := m-i-1, n-j-1
			if x == 0 && row0 {
				matrix[x][y] = 0
			} else if y == 0 && col0 {
				matrix[x][y] = 0
			} else if matrix[x][0] == 0 || matrix[0][y] == 0 {
				matrix[x][y] = 0
			}
		}
	}
}
