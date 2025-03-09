package p48

func swap(matrix [][]int, i, j int) {
	n := len(matrix)

	d1x, d1y := j, n-i
	d2x, d2y := n-i, n-j
	d3x, d3y := n-j, i

	temp := matrix[i][j]
	matrix[i][j] = matrix[d3x][d3y]
	matrix[d3x][d3y] = matrix[d2x][d2y]
	matrix[d2x][d2y] = matrix[d1x][d1y]
	matrix[d1x][d1y] = temp
}

func rotate(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	n := len(matrix)

	for i := range n / 2 {
		for j := range (n + 1) / 2 {
			swap(matrix, i, j)
		}
	}
}
