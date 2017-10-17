package p210

func FindOrder(numCourses int, prerequisites [][]int) (res []int) {
	res = make([]int, 0, numCourses)

	var graph = make([][]int, numCourses)
	var reversedGraph = make([]int, numCourses)

	for _, tuple := range prerequisites {
		var from, to = tuple[1], tuple[0]
		graph[from] = append(graph[from], to)
		reversedGraph[to] += 1
	}

	// node with no incoming edges.
	var s = make([]int, 0, numCourses)
	for i, l := range reversedGraph {
		if l == 0 {
			s = append(s, i)
		}
	}

	for len(s) > 0 {
		var from = s[0]
		s = s[1:]

		res = append(res, from)
		for _, to := range graph[from] {
			reversedGraph[to] -= 1
			if reversedGraph[to] == 0 {
				s = append(s, to)
			}
		}
	}

	for _, l := range reversedGraph {
		if l > 0 {
			res = make([]int, 0)
			return
		}
	}
	return
}
