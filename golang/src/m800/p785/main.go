package p785

func isBipartite(graph [][]int) bool {
	if len(graph) == 0 {
		return false
	}

	edgeMap := make(map[int]map[int]bool)
	for i, edges := range graph {
		if _, ok := edgeMap[i]; !ok {
			edgeMap[i] = make(map[int]bool)
		}
		for _, edge := range edges {
			edgeMap[i][edge] = true
		}
	}

	setA, setB := make(map[int]bool), make(map[int]bool)

	q, inA := []int{0}, true

	for len(q) > 0 {
		l := len(q)

		for i := range l {
			node := q[i]
			if inA {
				setA[node] = true
			} else {
				setB[node] = true
			}
			for next, _ := range edgeMap[node] {
				delete(edgeMap[next], node)
				q = append(q, next)
			}
		}

		q = q[l:]
		inA = !inA
	}

	return true
}
