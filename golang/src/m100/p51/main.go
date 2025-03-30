package p51

import "strings"

type Position struct {
	x, y int
}

func print(queens []Position) []string {
	n := len(queens)

	var res []string

	for _, queen := range queens {
		res = append(res, strings.Repeat(".", queen.y)+string('Q')+strings.Repeat(".", n-queen.y-1))
	}
	return res
}

func canPlace(queens []Position, x, y int) bool {
	for _, queen := range queens {
		if queen.y == y || queen.x+queen.y == x+y || queen.x-queen.y == x-y {
			return false
		}
	}
	return true
}

func dfs(res *[][]string, queens []Position, n int) {
	if len(queens) == n {
		*res = append(*res, print(queens))
		return
	}

	x := len(queens)
	for y := range n {
		if canPlace(queens, x, y) {
			queens = append(queens, Position{x, y})
			dfs(res, queens, n)
			queens = queens[:len(queens)-1]
		}
	}
}

func solveNQueens(n int) [][]string {
	res := make([][]string, 0)
	dfs(&res, []Position{}, n)
	return res
}
