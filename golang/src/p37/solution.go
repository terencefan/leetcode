package main

type Solver struct {
	graph [][]byte
}

func NewSolver(graph [][]byte) *Solver {
	return &Solver{graph}
}

func (s *Solver) candidates(x, y int) (r []byte) {
	var c byte
	var m = make(map[byte]bool)

	for i := 0; i < 9; i++ {
		c = s.graph[x][i]
		m[c] = true

		c = s.graph[i][y]
		m[c] = true
	}

	midx, midy := x / 3 * 3 + 1, y / 3 * 3 + 1
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			c = s.graph[midx + i][midy + j]
			m[c] = true
		}
	}

	for i := byte('1'); i <= byte('9'); i++ {
		if !m[i] {
			r = append(r, i)
		}
	}
	return
}

func next(x, y int) (int, int) {
	if y == 8 {
		return x + 1, 0
	} else {
		return x, y + 1
	}
}

func (s *Solver) dfs(x, y int) bool {
	if x == 9 && y == 0 {
		return true
	}

	if s.graph[x][y] == '.' {
		for _, k := range s.candidates(x, y){
			s.graph[x][y] = k
			if s.dfs(next(x, y)) {
				return true
			}
			s.graph[x][y] = '.'
		}
	} else {
		return s.dfs(next(x, y))
	}
	return false
}

func solveSudoku(board [][]byte)  {
	s := NewSolver(board)
	s.dfs(0, 0)
}

func main() {
	solveSudoku([][]byte{
		{'5','3','.','.','7','.','.','.','.'},
		{'6','.','.','1','9','5','.','.','.'},
		{'.','9','8','.','.','.','.','6','.'},
		{'8','.','.','.','6','.','.','.','3'},
		{'4','.','.','8','.','3','.','.','1'},
		{'7','.','.','.','2','.','.','.','6'},
		{'.','6','.','.','.','.','2','8','.'},
		{'.','.','.','4','1','9','.','.','5'},
		{'.','.','.','.','8','.','.','7','9'},
	})
}