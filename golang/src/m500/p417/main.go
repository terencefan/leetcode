package p417

const (
	PACIFIC  = 0x1
	ATLANTIC = 0x10
)

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func bfs(heights, visited [][]int, i, j, ocean int) {
	visited[i][j] |= ocean
	q := [][]int{{i, j}}

	for len(q) > 0 {
		x, y := q[0][0], q[0][1]
		for _, d := range directions {
			nx, ny := x+d[0], y+d[1]
			if nx >= 0 && nx < len(heights) && ny >= 0 && ny < len(heights[0]) && heights[nx][ny] >= heights[x][y] && visited[nx][ny]&ocean == 0 {
				visited[nx][ny] |= ocean
				q = append(q, []int{nx, ny})
			}
		}
		q = q[1:]
	}
}

func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 {
		return [][]int{}
	}
	m, n := len(heights), len(heights[0])

	visited := make([][]int, m)
	for i := range m {
		visited[i] = make([]int, n)
	}

	for i := range m {
		bfs(heights, visited, i, 0, PACIFIC)
		bfs(heights, visited, i, n-1, ATLANTIC)
	}
	for i := range n {
		bfs(heights, visited, 0, i, PACIFIC)
		bfs(heights, visited, m-1, i, ATLANTIC)
	}

	var r = make([][]int, 0)
	for i, row := range visited {
		for j, v := range row {
			if v == PACIFIC|ATLANTIC {
				r = append(r, []int{i, j})
			}
		}
	}
	return r
}
