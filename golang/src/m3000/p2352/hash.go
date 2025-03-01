package p2352

const MOD = 1e9 + 7

func test(grid [][]int, x, y int) bool {
	for i := range len(grid) {
		if grid[x][i] != grid[i][y] {
			return false
		}
	}
	return true
}

func equalPairs(grid [][]int) int {
	n := len(grid)
	xHash := make([]int, n)
	yHash := make([]int, n)

	for i := range n {
		x, y := 0, 0
		for j := range n {
			x *= 17
			x += grid[i][j]
			x %= MOD

			y *= 17
			y += grid[j][i]
			y %= MOD
		}
		xHash[i] = x
		yHash[i] = y
	}

	var r = 0
	for i := range n {
		for j := range n {
			if xHash[i] == yHash[j] && test(grid, i, j) {
				r++
			}
		}
	}
	return r
}
