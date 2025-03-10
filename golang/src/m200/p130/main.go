package p130

const (
	EMPTY byte = 'O'
	WALL  byte = 'X'
	PASS  byte = 'P'
)

var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func onEdge(board [][]byte, i, j int) bool {
	return i == 0 || j == 0 || i == len(board)-1 || j == len(board[0])-1
}

func inBound(board [][]byte, i, j int) bool {
	return i >= 0 && i < len(board) && j >= 0 && j < len(board[0])
}

func search(board [][]byte, i, j int) byte {
	if onEdge(board, i, j) {
		return EMPTY
	}

	for _, d := range directions {
		nx, ny := i+d[0], j+d[1]
		if inBound(board, nx, ny) && board[nx][ny] == EMPTY {
			board[nx][ny] = PASS
			if search(board, nx, ny) == EMPTY {
				return EMPTY
			}
		}
	}
	return WALL
}

func fill(board [][]byte, i, j int, f byte) {
	for _, d := range directions {
		nx, ny := i+d[0], j+d[1]
		if inBound(board, nx, ny) && board[nx][ny] != WALL {
			board[nx][ny] = f
			fill(board, nx, ny, f)
		}
	}
}

func solve(board [][]byte) {
	if len(board) == 0 || len(board[0]) == 0 {
		return
	}
	m, n := len(board), len(board[0])

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] != EMPTY {
				continue
			}
			board[i][j] = PASS
			if search(board, i, j) == WALL {
				board[i][j] = WALL
				fill(board, i, j, WALL)
			} else {
				fill(board, i, j, PASS)
			}
		}
	}

	for i := range m {
		for j := range n {
			if board[i][j] == PASS {
				board[i][j] = EMPTY
			}
		}
	}
}
