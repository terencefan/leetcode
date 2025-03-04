package p74

func to(x, m int) (int, int) {
	return x / m, x % m
}

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}

	m, n := len(matrix), len(matrix[0])

	l, r := 0, m*n

	for l < r {
		mid := l + (r-l)/2

		i, j := to(mid, m)
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return false
}
