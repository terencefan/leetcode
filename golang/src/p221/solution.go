package main

import (
	"fmt"
)

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

type Matrix [][]byte

func (m *Matrix) get(i, j int) int {
	if i < 0 || j < 0 {
		return 0
	} else if i > len(*m) {
		return 0
	} else if j > len((*m)[0]) {
		return 0
	} else {
		return int((*m)[i][j])
	}
}

func (m *Matrix) update(i, j int) int {
	x := m.get(i - 1, j - 1)
	x = min(x, m.get(i, j - 1))
	x = min(x, m.get(i - 1, j))
	if (*m)[i][j] != 0 {
		(*m)[i][j] = byte(x + 1)
		return x + 1
	}
	return 0
}

func maximalSquare(matrix Matrix) int {
	if len(matrix) == 0 {
		return 0
	} else if len(matrix[0]) == 0 {
		return 0
	}

	var r = 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[i][j] -= '0'

			t := matrix.update(i, j)
			if t > r {
				r = t
			}
		}
	}
	return r * r
}

func main() {
	r := maximalSquare([][]byte{
		{'1', '1'},
		{'1', '1'},
	})
	fmt.Println(r)
}
