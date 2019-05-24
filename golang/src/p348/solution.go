package main

type TicTacToe struct {
	board [][]int
}


/** Initialize your data structure here. */
func Constructor(n int) TicTacToe {
	board := make([][]int, n)
	for i := 0; i < n; i++ {
		board[i] = make([]int, n)
	}
	return TicTacToe{
		board: board,
	}
}


/** Player {player} makes a move at ({row}, {col}).
        @param row The row of the board.
        @param col The column of the board.
        @param player The player, can be either 1 or 2.
        @return The current winning condition, can be either:
                0: No one wins.
                1: Player 1 wins.
                2: Player 2 wins. */
func (t *TicTacToe) Move(row int, col int, player int) int {
	t.board[row][col] = player
	c1, c2, c3, c4 := true, true, true, true
	for i := 0; i < len(t.board); i++ {
		c1 = c1 && t.board[row][i] == player
		c2 = c2 && t.board[i][col] == player
		c3 = c3 && t.board[i][i] == player
		c4 = c4 && t.board[len(t.board) - 1][i] == player
	}
	if c1 || c2 || c3 || c4 {
		return player
	} else {
		return 0
	}
}


func main() {
	obj := Constructor(2);
	obj.Move(0, 1, 1)
	obj.Move(0, 1, 1)
	obj.Move(0, 1, 1)
}
