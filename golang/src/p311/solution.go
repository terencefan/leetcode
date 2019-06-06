package main

type Rows map[int]int
type Cols map[int]int

type SparseMatrix struct {
	height, width int
	rows          map[int]Cols
	cols          map[int]Rows
}

func Construct(matrix [][]int) SparseMatrix {
	rows := make(map[int]Cols)
	cols := make(map[int]Rows)

	width := 0
	if len(matrix) > 0 {
		width = len(matrix[0])
	}

	for row, columns := range matrix {
		for col, val := range columns {
			if rows[row] == nil {
				rows[row] = make(Cols)
			}
			if cols[col] == nil {
				cols[col] = make(Rows)
			}
			rows[row][col] = val
			cols[col][row] = val
		}
	}
	return SparseMatrix{
		height: len(matrix),
		width:  width,
		rows:   rows,
		cols:   cols,
	}
}

func deal(a, b map[int]int) (r int) {
	for x, v := range a {
		r += b[x] * v
	}
	return
}

func multi(a, b SparseMatrix) (r [][]int) {
	if a.width != b.height {
		return [][]int{}
	}

	width, height := b.width, a.height

	r = make([][]int, height)
	for i := 0; i < height; i++ {
		r[i] = make([]int, width)
		for j := 0; j < width; j++ {
			r[i][j] = deal(a.rows[i], b.cols[j])
		}
	}
	return
}

func multiply(A [][]int, B [][]int) [][]int {
	return multi(Construct(A), Construct(B))
}
