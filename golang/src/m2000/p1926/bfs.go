package p1926

const Wall = '+'
const Empty = '.'

func IsEdge(maze [][]byte, x, y int) bool {
	return (x == 0 || y == 0 || x == len(maze)-1 || y == len(maze[0])-1)
}

func CanVisit(maze [][]byte, x, y int) bool {
	return x >= 0 && x < len(maze) && y >= 0 && y < len(maze[x]) && maze[x][y] == Empty
}

func ShouldVisit(grid [][]int, x, y, step int) bool {
	return grid[x][y] == 0 || grid[x][y] > step
}

var Directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func nearestExit(maze [][]byte, entrance []int) int {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return -1
	}
	m, n := len(maze), len(maze[0])

	var stack = make([][]int, 0)

	stack = append(stack, []int{entrance[0], entrance[1], 0})
	maze[entrance[0]][entrance[1]] = Wall

	var grid = make([][]int, len(maze))
	for i := range maze {
		grid[i] = make([]int, len(maze[i]))
	}

	for len(stack) > 0 {
		var item = stack[len(stack)-1]
		var x, y, step = item[0], item[1], item[2]
		stack = stack[:len(stack)-1]

		for _, d := range Directions {
			x1, y1 := x+d[0], y+d[1]
			if CanVisit(maze, x1, y1) && ShouldVisit(grid, x1, y1, step+1) {
				grid[x1][y1] = step + 1
				stack = append(stack, []int{x1, y1, step + 1})
			}
		}
	}

	var r = 1 << 31

	for i := range m {
		if grid[i][0] > 0 {
			r = min(r, grid[i][0])
		}
		if grid[i][n-1] > 0 {
			r = min(r, grid[i][n-1])
		}
	}

	for i := range n {
		if grid[0][i] > 0 {
			r = min(r, grid[0][i])
		}
		if grid[m-1][i] > 0 {
			r = min(r, grid[m-1][i])
		}
	}

	if r == 1<<31 {
		return -1
	} else {
		return r
	}
}
