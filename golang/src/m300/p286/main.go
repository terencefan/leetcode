package p286

const (
	WALL = -1
	GATE = 0
)

type Point struct {
	x, y int
}

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func wallsAndGates(rooms [][]int) {
	if len(rooms) == 0 {
		return
	}
	m, n := len(rooms), len(rooms[0])

	q := make([]Point, 0)
	for x, row := range rooms {
		for y, val := range row {
			if val == GATE {
				q = append(q, Point{x, y})
			}
		}
	}

	distance := 0
	for len(q) > 0 {
		distance++
		l := len(q)

		for i := range l {
			p := q[i]
			for _, d := range directions {
				nx, ny := p.x+d[0], p.y+d[1]
				if nx < 0 || ny < 0 || nx >= m || ny >= n || rooms[nx][ny] == WALL || rooms[nx][ny] <= distance {
					continue
				}
				rooms[nx][ny] = distance
				q = append(q, Point{nx, ny})
			}
		}
		q = q[l:]
	}
}
