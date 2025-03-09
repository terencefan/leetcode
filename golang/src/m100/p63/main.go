package p63

const OBSTABLE = 1

func uniquePathsWithObstacles(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	} else if len(grid[0]) == 0 {
		return 0
	} else if grid[0][0] == OBSTABLE {
		return 0
	}
	var dp = make([]int, len(grid[0]))
	for i, row := range grid {
		for j, val := range row {
			if val == OBSTABLE {
				dp[j] = 0
				continue
			} else if i == 0 && j == 0 {
				dp[j] = 1
				continue
			}
			if i > 0 && grid[i-1][j] == OBSTABLE {
				dp[j] = 0
			}
			if j > 0 && grid[i][j-1] != OBSTABLE {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[len(grid[0])-1]
}
