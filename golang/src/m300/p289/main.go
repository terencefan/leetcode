package p289

const (
	DEAD             = 0
	LIVE             = 1
	UNDER_POPULATION = 2
	NEXT_GENERATION  = 3
	OVER_POPULATION  = 4
	REPDODUCTION     = 5
)

var directions = [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}, {-1, -1}, {-1, 1}, {1, -1}, {1, 1}}

func update(board [][]int, i, j int) {
	live := 0
	for _, d := range directions {
		x, y := i+d[0], j+d[1]
		if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
			continue
		}
		switch board[x][y] {
		case LIVE, UNDER_POPULATION, NEXT_GENERATION, OVER_POPULATION:
			live++
		}
	}

	if board[i][j] == LIVE {
		if live < 2 {
			board[i][j] = UNDER_POPULATION
		} else if live == 2 || live == 3 {
			board[i][j] = LIVE
		} else {
			board[i][j] = OVER_POPULATION
		}
	} else {
		if live == 3 {
			board[i][j] = REPDODUCTION
		}
	}
}

func gameOfLife(board [][]int) {
	for i, row := range board {
		for j := range row {
			update(board, i, j)
		}
	}
	for i, row := range board {
		for j := range row {
			switch board[i][j] {
			case UNDER_POPULATION, OVER_POPULATION:
				board[i][j] = DEAD
			case NEXT_GENERATION, REPDODUCTION:
				board[i][j] = LIVE
			}
		}
	}
}
