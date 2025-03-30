package p473

import (
	"slices"
	"sort"
)

type Square struct {
	a, b, c, d int
}

func (s Square) Simplify() Square {
	values := []int{s.a, s.b, s.c, s.d}
	sort.Ints(values)
	return Square{values[0], values[1], values[2], values[3]}
}

func dfs(matchsticks []int, visited map[Square]bool, s Square, n int) bool {
	if len(matchsticks) == 0 {
		return true
	}

	s = s.Simplify()
	if visited[s] {
		return false
	}

	stick := matchsticks[0]

	if s.a+stick <= n {
		if dfs(matchsticks[1:], visited, Square{s.a + stick, s.b, s.c, s.d}, n) {
			return true
		}
	}
	if s.b+stick <= n {
		if dfs(matchsticks[1:], visited, Square{s.a, s.b + stick, s.c, s.d}, n) {
			return true
		}
	}
	if s.c+stick <= n {
		if dfs(matchsticks[1:], visited, Square{s.a, s.b, s.c + stick, s.d}, n) {
			return true
		}
	}
	if s.d+stick <= n {
		if dfs(matchsticks[1:], visited, Square{s.a, s.b, s.c, s.d + stick}, n) {
			return true
		}
	}

	visited[s] = true
	return false
}

func makesquare(matchsticks []int) bool {

	sum := 0
	for _, stick := range matchsticks {
		sum += stick
	}
	slices.Sort(matchsticks)
	slices.Reverse(matchsticks)

	if sum%4 != 0 {
		return false
	}
	var visited = make(map[Square]bool)
	return dfs(matchsticks, visited, Square{0, 0, 0, 0}, sum/4)
}
