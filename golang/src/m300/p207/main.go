package p207

func canFinish(numCourses int, prerequisites [][]int) bool {
	var requires = make([]int, numCourses)
	var depends = make([][]int, numCourses)

	for _, item := range prerequisites {
		a, b := item[0], item[1]
		depends[a] = append(depends[a], b)
		requires[b]++
	}

	q := make([]int, 0)
	for i, req := range requires {
		if req == 0 {
			q = append(q, i)
		}
	}

	for len(q) > 0 {
		course := q[0]
		q = q[1:]

		for _, next := range depends[course] {
			requires[next]--
			if requires[next] == 0 {
				q = append(q, next)
			}
		}
	}

	for _, req := range requires {
		if req > 0 {
			return false
		}
	}
	return true
}
