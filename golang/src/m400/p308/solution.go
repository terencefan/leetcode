package main

type NumMatrix struct {
	origin [][]int
	sum    [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{matrix, make([][]int, 0)}
	}

	sum := make([][]int, len(matrix))
	for i, row := range matrix {
		sum[i] = make([]int, len(matrix[0]))
		s := 0
		for j, val := range row {
			s += val
			sum[i][j] = s
		}
	}
	return NumMatrix{matrix, sum}
}

func (m *NumMatrix) Update(row int, col int, val int) {
	delta := val - m.origin[row][col]
	m.origin[row][col] = val

	for i := col; i < len(m.sum[row]); i++ {
		m.sum[row][i] += delta
	}
}

func (m *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) (r int) {
	if row1 > row2 {
		row1, row2 = row2, row1
	}
	if col1 > col2 {
		col1, col2 = col2, col1
	}

	for i := row1; i <= row2; i++ {
		r += m.sum[i][col2]
		if col1 > 0 {
			r -= m.sum[i][col1-1]
		}
	}
	return r
}
