package p36

const EMPTY = '.'

type Sodoku [9]bool

func (s *Sodoku) add(num int) bool {
	if (*s)[num-1] {
		return false
	}
	(*s)[num-1] = true
	return true
}

func isValidSudoku(board [][]byte) bool {
	if len(board) != 9 || len(board[0]) != 9 {
		return false
	}

	var rows = make([]Sodoku, 9)
	var cols = make([]Sodoku, 9)
	var blocks = make([]Sodoku, 9)

	for i := range 9 {
		for j := range 9 {
			if board[i][j] == EMPTY {
				continue
			}
			num := int(board[i][j] - '0')

			b := i/3*3 + j/3

			if !rows[i].add(num) || !cols[j].add(num) || !blocks[b].add(num) {
				return false
			}
		}
	}
	return true
}
