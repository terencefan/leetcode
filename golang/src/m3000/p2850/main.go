package p2850

type Grid [][]int

func hash(grid Grid) int {
	var res = 0
	for _, row := range grid {
		for _, num := range row {
			res = res*10 + num
		}
	}
	return res
}

func dfs(grid Grid, visited map[int]int, steps int) {
	for i, row := range grid {
		for j, val := range row {
			if val <= 1 {
				continue
			}
			for _, d := range directions {
				nx, ny := i+d[0], j+d[1]
				if nx >= 0 && ny >= 0 && nx < 3 && ny < 3 && grid[nx][ny] < grid[i][j] {
					grid[nx][ny]++
					grid[i][j]--
					key := hash(grid)
					if visited[key] == 0 || visited[key] > steps+1 {
						visited[key] = steps + 1
						dfs(grid, visited, steps+1)
					}
					grid[nx][ny]--
					grid[i][j]++
				}
			}
		}
	}
}

var directions = [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func minimumMoves(grid Grid) int {
	visited := make(map[int]int)
	dfs(grid, visited, 0)
	return visited[111111111]
}
