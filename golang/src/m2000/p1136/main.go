package p1136

func minimumSemesters(n int, relations [][]int) int {
	var deps = make(map[int][]int)
	var inDegree = make([]int, n+1)

	for _, r := range relations {
		deps[r[0]] = append(deps[r[0]], r[1])
		inDegree[r[1]]++
	}

	q := make([]int, 0)
	for i := 1; i <= n; i++ {
		if inDegree[i] == 0 {
			q = append(q, i)
		}
	}

	semesters, courses := 0, 0
	for len(q) > 0 {
		l := len(q)
		courses += l

		for i := range l {
			course := q[i]
			for _, next := range deps[course] {
				inDegree[next]--
				if inDegree[next] == 0 {
					q = append(q, next)
				}
			}
		}

		q = q[l:]
		semesters++
	}

	if courses != n {
		return -1
	}
	return semesters
}
