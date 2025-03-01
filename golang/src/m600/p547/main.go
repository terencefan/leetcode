package p547

func dfs(isConnected [][]int, visited map[int]bool, city int) {
	if visited[city] {
		return
	}
	visited[city] = true

	for nextCity, connected := range isConnected[city] {
		if connected == 1 {
			dfs(isConnected, visited, nextCity)
		}
	}
}

func findCircleNum(isConnected [][]int) int {
	var r = 0
	var visited = map[int]bool{}

	for city := range isConnected {
		if visited[city] {
			continue
		}
		dfs(isConnected, visited, city)
		r++
	}
	return r
}
