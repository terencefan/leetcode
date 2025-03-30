package p269

func alienOrder(words []string) string {
	var edges = make(map[byte]map[byte]int, 0)
	var indegree = make(map[byte]int, 0)

	for _, word := range words {
		for _, c := range word {
			indegree[byte(c)] = 0
		}
	}

	for i := range len(words) - 1 {
		prev, next := words[i], words[i+1]

		k := 0
		for ; k < len(prev) && k < len(next); k++ {
			from, to := prev[k], next[k]
			if from != to {
				if _, ok := edges[from]; !ok {
					edges[from] = make(map[byte]int, 0)
				}
				if edges[from][to] == 0 {
					edges[from][to]++
					indegree[to]++
				}
				break
			}
		}
		if k == len(next) && len(prev) > len(next) {
			return ""
		}
	}

	q := make([]byte, 0)
	for c, d := range indegree {
		if d == 0 {
			q = append(q, c)
		}
	}

	var orders = make([]byte, 0)

	for len(q) > 0 {
		c := q[0]
		q = q[1:]
		orders = append(orders, c)

		for next, _ := range edges[c] {
			indegree[next]--
			if indegree[next] == 0 {
				q = append(q, next)
			}
		}
	}

	for _, c := range indegree {
		if c != 0 {
			return ""
		}
	}
	return string(orders)
}
