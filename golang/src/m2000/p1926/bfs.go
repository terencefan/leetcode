package p1926

const Wall = '+'
const Visited = '-'
const Empty = '.'

func IsEdge(maze [][]byte, x, y int) bool {
	return (x == 0 || y == 0 || x == len(maze)-1 || y == len(maze[0])-1)
}

func CanVisit(maze [][]byte, x, y int) bool {
	return x >= 0 && x < len(maze) && y >= 0 && y < len(maze[x]) && maze[x][y] == Empty
}

var Directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type Point struct {
	x, y, step int
}

func nearestExit(maze [][]byte, entrance []int) int {
	if len(maze) == 0 || len(maze[0]) == 0 {
		return -1
	}

	var queue = []Point{{entrance[0], entrance[1], 0}}
	maze[entrance[0]][entrance[1]] = Wall

	for len(queue) > 0 {
		point := queue[0]
		queue = queue[1:]

		if IsEdge(maze, point.x, point.y) && maze[point.x][point.y] != Wall {
			return point.step
		}

		for _, d := range Directions {
			dx, dy := point.x+d[0], point.y+d[1]
			if CanVisit(maze, dx, dy) {
				queue = append(queue, Point{dx, dy, point.step + 1})
				maze[dx][dy] = Visited
			}
		}
	}
	return -1
}
