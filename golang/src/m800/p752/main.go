package p752

type lock struct {
	a, b, c, d int
}

func next(l lock) []lock {
	var res []lock
	res = append(res, lock{l.a, l.b, l.c, (l.d + 1) % 10})
	res = append(res, lock{l.a, l.b, l.c, (l.d + 9) % 10})
	res = append(res, lock{l.a, l.b, (l.c + 1) % 10, l.d})
	res = append(res, lock{l.a, l.b, (l.c + 9) % 10, l.d})
	res = append(res, lock{l.a, (l.b + 1) % 10, l.c, l.d})
	res = append(res, lock{l.a, (l.b + 9) % 10, l.c, l.d})
	res = append(res, lock{(l.a + 1) % 10, l.b, l.c, l.d})
	res = append(res, lock{(l.a + 9) % 10, l.b, l.c, l.d})
	return res
}

func str2lock(s string) lock {
	return lock{int(s[0] - '0'), int(s[1] - '0'), int(s[2] - '0'), int(s[3] - '0')}
}

func openLock(deadends []string, target string) int {
	var visited = make(map[lock]bool)

	for _, deadend := range deadends {
		visited[str2lock(deadend)] = true
	}
	var start = lock{0, 0, 0, 0}
	if visited[start] {
		return -1
	}

	var end = str2lock(target)
	visited[start] = true

	q := []lock{start}

	steps := 0
	for len(q) > 0 {
		length := len(q)

		for i := range length {
			current := q[i]
			if current == end {
				return steps
			}
			for _, next := range next(current) {
				if !visited[next] {
					visited[next] = true
					q = append(q, next)
				}
			}
		}
		q = q[length:]
		steps++
	}
	return -1
}
