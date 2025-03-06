package p210

func findOrder(numCourses int, prerequisites [][]int) []int {
	var inDegrees = make([]int, numCourses)

	var deps = make(map[int][]int)

	for _, p := range prerequisites {
		b, a := p[0], p[1]
		inDegrees[b]++
		deps[a] = append(deps[a], b)
	}

	q := make([]int, 0)
	for i, v := range inDegrees {
		if v == 0 {
			q = append(q, i)
		}
	}

	var r = make([]int, 0)
	for len(q) > 0 {
		l := len(q)

		for i := range l {
			for _, b := range deps[q[i]] {
				inDegrees[b]--
				if inDegrees[b] == 0 {
					q = append(q, b)
				}
			}
		}
		r = append(r, q[:l]...)
		q = q[l:]
	}

	if len(r) != numCourses {
		return []int{}
	}
	return r
}
