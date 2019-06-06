package main

import (
	"fmt"
)

func tryRemove(board []string, char byte) bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[i][j] == char {
				t := board[i]
				board[i] = board[i][:j] + " " + board[i][j+1:]
				if !isEndingState(board, 'O') && !isEndingState(board, 'X') {
					return true
				}
				board[i] = t
			}
		}
	}
	return false
}

func isEndingState(board []string, char byte) bool {
	tl2br, tr2bl := true, true
	for i := 0; i < 3; i++ {
		row, col := true, true
		for j := 0; j < 3; j++ {
			if board[i][j] != char {
				row = false
			}
			if board[j][i] != char {
				col = false
			}
		}
		if row || col {
			return true
		}

		if board[i][i] != char {
			tl2br = false
		}
		if board[i][2-i] != char {
			tr2bl = false
		}
	}
	return tl2br || tr2bl
}

func validTicTacToe(board []string) bool {
	if len(board) != 3 {
		return false
	}

	countX, countO := 0, 0
	for _, row := range board {
		if len(row) != 3 {
			return false
		}
		for _, char := range row {
			switch char {
			case 'X':
				countX++
			case 'O':
				countO++
			case ' ':
				continue
			default:
				return false
			}
		}
	}
	if countO > countX || countX > countO+1 {
		return false
	}

	if countX == 0 {
		return true
	}

	if !isEndingState(board, 'O') && !isEndingState(board, 'X') {
		return true
	}

	if countO == countX {
		return tryRemove(board, 'O')
		// check if remove a 'O' can reach the non-ending state.
	} else {
		return tryRemove(board, 'X')
		// check if remove a 'X' can reach the non-ending state.
	}
}

func main() {
	r := validTicTacToe([]string{
		"XXX",
		"XOO",
		"OO ",
	})
	fmt.Println(r)
}
