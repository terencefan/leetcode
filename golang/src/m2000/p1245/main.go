package p1245

func treeDiameter(edges [][]int) int {
	var paths = make(map[int][]int)

	for _, edge := range edges {
		start, end := edge[0], edge[1]
		paths[start] = append(paths[start], end)
	}

	return 0
}
