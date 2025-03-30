package p130

const (
	EMPTY byte = 'O'
	WALL  byte = 'X'
	LIVE  byte = 'V'
)

var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func onEdge(board [][]byte, i, j int) bool {
	return i == 0 || j == 0 || i == len(board)-1 || j == len(board[0])-1
}

func inBound(board [][]byte, i, j int) bool {
	return i >= 0 && i < len(board) && j >= 0 && j < len(board[0])
}

func fill(board [][]byte, i, j int) {
	for _, d := range directions {
		x, y := i+d[0], j+d[1]
		if inBound(board, x, y) && board[x][y] == EMPTY {
			board[x][y] = LIVE
			fill(board, x, y)
		}
	}
}

func solve(board [][]byte) {
	for i, row := range board {
		for j, cell := range row {
			if cell == EMPTY && onEdge(board, i, j) {
				board[i][j] = LIVE
				fill(board, i, j)
			}
		}
	}

	for i, row := range board {
		for j, cell := range row {
			if cell == EMPTY {
				board[i][j] = WALL
			} else if cell == LIVE {
				board[i][j] = EMPTY
			}
		}
	}
}
