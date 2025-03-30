package p815

func numBusesToDestination(routes [][]int, source int, target int) int {
	if source == target {
		return 0
	}

	stopRoutes := make(map[int][]int)
	targetRoutes := make(map[int]bool)
	routeMap := make([]map[int]bool, len(routes))
	for i := range len(routes) {
		routeMap[i] = make(map[int]bool)
	}

	for i, stops := range routes {
		for _, stop := range stops {
			if stop == target {
				targetRoutes[i] = true
			}
			for _, k := range stopRoutes[stop] {
				routeMap[i][k] = true
				routeMap[k][i] = true
			}
			stopRoutes[stop] = append(stopRoutes[stop], i)
		}
	}

	if len(targetRoutes) == 0 {
		return -1
	}

	q := []int{}
	visited := make(map[int]bool)
	for _, route := range stopRoutes[source] {
		if targetRoutes[route] {
			return 1
		}
		q = append(q, route)
		visited[route] = true
	}

	steps := 2
	for len(q) > 0 {
		l := len(q)

		for i := range l {
			route := q[i]
			for next := range routeMap[route] {
				if targetRoutes[next] {
					return steps
				} else if !visited[next] {
					q = append(q, next)
					visited[next] = true
				}
			}
		}
		q = q[l:]
		steps++
	}
	return -1
}
