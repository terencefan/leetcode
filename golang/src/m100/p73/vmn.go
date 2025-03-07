package p73

func setZeroes1(matrix [][]int) {
	if len(matrix) == 0 {
		return
	}
	m, n := len(matrix), len(matrix[0])

	var m1 = make([]bool, m)
	var n1 = make([]bool, n)

	for i := range m {
		for j := range n {
			if matrix[i][j] == 0 {
				m1[i] = true
				n1[j] = true
			}
		}
	}

	for i, v := range m1 {
		if v {
			for j := range n {
				matrix[i][j] = 0
			}
		}
	}

	for j, v := range n1 {
		if v {
			for i := range m {
				matrix[i][j] = 0
			}
		}
	}
}
