package main

type point struct {
	x, y int
}

const land byte = '1'

var arr = []int{1, 0, -1, 0}
var height, width int

func dfs(grid *[][]byte, p *point) {
	stack := make([]*point, 0, height+width)
	stack = append(stack, p)

	for length := 0; length >= 0; {
		cur := stack[length]
		length--
		(*grid)[cur.x][cur.y] = 0

		for i := 0; i < 4; i++ {
			dx, dy := arr[i], arr[(i+1)%4]
			nx, ny := cur.x+dx, cur.y+dy

			if nx < 0 || nx >= height {
				continue
			} else if ny < 0 || ny >= width {
				continue
			} else if (*grid)[nx][ny] != land {
				continue
			} else {
				length++
			}

			p := &point{dx + cur.x, dy + cur.y}
			if len(stack) == length {
				stack = append(stack, p)
			} else {
				stack[length] = p
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	height, width = len(grid), len((grid)[0])

	r := 0
	for x := 0; x < height; x++ {
		for y := 0; y < width; y++ {
			if grid[x][y] == land {
				dfs(&grid, &point{x, y})
				r++
			}
		}
	}
	return r
}
