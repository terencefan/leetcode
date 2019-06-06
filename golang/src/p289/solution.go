package main

import (
	"fmt"
)

type Board [][]int

func (b Board) get(x, y int) int {
	if x < 0 || y < 0 {
		return 0
	} else if x > len(b)-1 {
		return 0
	} else if y > len(b[0])-1 {
		return 0
	} else {
		return b[x][y] % 2
	}
}

func (b Board) adjust(x, y int) int {
	lives := 0
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i == x && j == y {
				continue
			}
			lives += b.get(i, j)
		}
	}
	if b[x][y] == 0 {
		if lives == 3 {
			return 1
		}
	} else {
		if lives >= 2 && lives <= 3 {
			return 1
		}
	}
	return 0
}

func gameOfLife(board Board) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			next := board.adjust(i, j)
			board[i][j] |= next << 1
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			board[i][j] >>= 1
		}
	}
}

func main() {
	var b = [][]int{
		{0, 1, 0},
		{0, 0, 1},
		{1, 1, 1},
		{0, 0, 0},
	}
	gameOfLife(b)
	fmt.Println(b)
}
