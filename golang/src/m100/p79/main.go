package p79

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

func search(board [][]byte, word string, i, j, k int) bool {
	if board[i][j] != word[k] {
		return false
	} else if k == len(word)-1 {
		return true
	}

	board[i][j] = '.'
	for _, d := range directions {
		x, y := i+d[0], j+d[1]
		if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
			continue // out of border
		}
		if search(board, word, x, y, k+1) {
			return true
		}
	}
	board[i][j] = word[k]
	return false
}

func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	} else if len(word) == 0 {
		return true
	}

	for i, row := range board {
		for j := range row {
			if search(board, word, i, j, 0) {
				return true
			}
		}
	}
	return false
}
