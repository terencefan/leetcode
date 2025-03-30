package p785

func isBipartite(graph [][]int) bool {
	if len(graph) == 0 {
		return false
	}

	edgeMap := make([]map[int]bool, len(graph))
	for i := range len(graph) {
		edgeMap[i] = make(map[int]bool)
	}

	for from, nodes := range graph {
		for _, node := range nodes {
			edgeMap[from][node] = true
			edgeMap[node][from] = true
		}
	}

	setA, setB := make(map[int]bool), make(map[int]bool)
	for k := range graph {
		if setA[k] || setB[k] {
			continue
		}
		q, inA := []int{k}, true

		for len(q) > 0 {
			l := len(q)

			for i := range l {
				node := q[i]
				if inA {
					if setB[node] {
						return false
					}
					setA[node] = true
				} else {
					if setA[node] {
						return false
					}
					setB[node] = true
				}
				for next := range edgeMap[node] {
					delete(edgeMap[next], node)
					delete(edgeMap[node], next)
					q = append(q, next)
				}
			}

			q = q[l:]
			inA = !inA
		}
	}
	return true
}
