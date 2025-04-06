package p1245

func treeDiameter(edges [][]int) int {
	var paths = make(map[int][]int)

	for _, edge := range edges {
		start, end := edge[0], edge[1]
		paths[start] = append(paths[start], end)
		paths[end] = append(paths[end], start)
	}

	visited := map[int]bool{0: true}

	q := []int{0}
	for len(q) > 0 {
		l := len(q)

		for i := range l {
			node := q[i]
			for _, next := range paths[node] {
				if !visited[next] {
					q = append(q, next)
					visited[next] = true
				}
			}
		}
		if l == len(q) {
			break
		}
		q = q[l:]
	}

	q = q[:1]
	visited = map[int]bool{q[0]: true}
	distance := 0

	for len(q) > 0 {
		l := len(q)

		for i := range l {
			node := q[i]
			for _, next := range paths[node] {
				if !visited[next] {
					q = append(q, next)
					visited[next] = true
				}
			}
		}
		q = q[l:]
		distance++
	}
	return distance - 1
}
