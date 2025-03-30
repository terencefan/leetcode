package p323

func countComponents(n int, edges [][]int) int {

	visited := make(map[int]bool)

	edgeMap := make(map[int][]int)
	for _, edge := range edges {
		edgeMap[edge[0]] = append(edgeMap[edge[0]], edge[1])
		edgeMap[edge[1]] = append(edgeMap[edge[1]], edge[0])
	}

	components := 0
	for i := range n {
		if visited[i] {
			continue
		}
		components++

		visited[i] = true
		q := []int{i}
		for len(q) > 0 {
			for _, next := range edgeMap[q[0]] {
				if !visited[next] {
					q = append(q, next)
					visited[next] = true
				}
			}
			q = q[1:]
		}
	}
	return components
}
